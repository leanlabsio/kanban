package gitlab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

type ResponseError struct {
	message string
	StatusCode int
}

// Type response error for usage logic
func (r ResponseError) Error() string {
	return r.message
}

// Do sends an HTTP request and returns an HTTP response, following
// policy (e.g. redirects, cookies, auth) as configured on the client.
func (g *GitlabContext) Do(req *http.Request, v interface{}) (*http.Response, error) {
	res, err := g.client.Do(req)

	if err != nil {
		return res, err
	}

	err = CheckResponse(res)
	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return res, err
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		return res, err
	}

	return res, nil
}

// CheckResponse checks the API response for errors, and returns them if
// present.  A response is considered an error if it has a status code outside
// the 200 range.  API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse.  Any other
// response body will be silently ignored.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	var res interface{}
	json.NewDecoder(r.Body).Decode(&res)

	logs := fmt.Sprintf("Bad response code: %d \n Request url: %s \n Data %+v", r.StatusCode, r.Request.URL.RequestURI(), res)
	fmt.Print(logs)
	return ResponseError{
		message: logs,
		StatusCode: r.StatusCode,
	}
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash.  If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *GitlabContext) NewRequest(method string, urlStr []string, body interface{}) (*http.Request, error) {
	u := getUrl(urlStr)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

// GetUrl creates url with base host and path
func getUrl(p []string) string {
	return cfg.BasePath + "/" + strings.Join(p, "/")
}

// addOptions adds the parameters in opt as URL query parameters to s.  opt
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}
