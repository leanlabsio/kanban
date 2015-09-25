package gitlab

import (
	"net/http"
	"net/url"
)

// Label represents a GitLab label.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/labels.html
type Label struct {
	Color string `json:"color"`
	Name  string `json:"name"`
}

// ListLabels gets all labels for given project.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/labels.html#list-labels
func (g *GitlabContext) ListLabels(project_id string, o *ListOptions) ([]*Label, error) {
	path := getUrl([]string{"projects", url.QueryEscape(project_id), "labels"})
	u, err := addOptions(path, o)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("GET", u, nil)

	var ret []*Label
	if _, err := g.Do(req, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}
