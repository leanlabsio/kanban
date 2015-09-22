package gitlab

import (
	"net/http"
	"net/url"
)

type Issue struct {
	Assignee    *User    `json:"assignee"`
	Author      *User    `json:"author"`
	Description string   `json:"description"`
	Id          int64    `json:"id"`
	Iid         int64    `json:"iid"`
	Labels      []string `json:"labels"`
	ProjectId   int64    `json:"project_id"`
	State       string   `json:"state"`
	Title       string   `json:"title"`
}

type IssueListOptions struct {
	// State filters issues based on their state.  Possible values are: open,
	// closed.  Default is "open".
	State string `url:"state,omitempty"`

	ListOptions
}

// Get list issues for gitlab projects
func (g *GitlabContext) ListIssues(project_id string, o *IssueListOptions) ([]*Issue, error) {
	cl := g.client
	path := getUrl([]string{"projects", url.QueryEscape(project_id), "issues"})
	u, err := addOptions(path, o)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("GET", u, nil)

	var ret []*Issue
	if err := g.Do(cl, req, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}
