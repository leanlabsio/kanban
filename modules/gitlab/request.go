package gitlab

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

// Execute request to gitlab and chek err
func (g *GitlabContext) Do(cl *http.Client, req *http.Request, v interface{}) error {
	res, err := cl.Do(req)

	if err != nil {
		log.Printf("%s", err.Error())
		return err
	}

	if res.StatusCode != http.StatusOK {
		log.Printf("%s: %i", "Bad response code", res.StatusCode)
		return errors.New(fmt.Sprintf("%s: %i", "Bad response code", res.StatusCode))
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		log.Printf("%s: %i", "Bad response code", err.Error())
		return err
	}

	return nil
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
