package gitlab

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
)

// File represents uploaded file to gitlab
//
// Gitlab API docs: http://docs.gitlab.com/ee/api/projects.html#upload-a-file
type File struct {
	Alt      string `json:"alt"`
	URL      string `json:"url"`
	Markdown string `json:"markdown"`
}

// UploadFile uploads file to gitlab project
//
// Gitlab API docs: http://docs.gitlab.com/ee/api/projects.html#upload-a-file
func (g *GitlabContext) UploadFile(projectID, name string, file io.Reader) (*File, error) {
	path := getUrl([]string{"projects", projectID, "uploads"})

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", name)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("POST", path, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	var ret File
	if _, err := g.Do(req, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
