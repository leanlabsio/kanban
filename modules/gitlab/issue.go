package gitlab

import (
	"net/http"
	"net/url"
)

type Issue struct {
	Assignee    *User      `json:"assignee"`
	Author      *User      `json:"author"`
	Description string     `json:"description"`
	Milestone   *Milestone `json:"milestone"`
	Id          int64      `json:"id"`
	Iid         int64      `json:"iid"`
	Labels      []string   `json:"labels"`
	ProjectId   int64      `json:"project_id"`
	State       string     `json:"state"`
	Title       string     `json:"title"`
}

type IssueRequest struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	AssigneeId  int64  `json:"assignee_id,omitempty"`
	MilestoneId int64  `json:"milestone_id,omitempty"`
	Labels      string `json:"labels,omitempty"`
	StateEvent  string `json:"state_event,omitempty"`
}

type IssueListOptions struct {
	// State filters issues based on their state.  Possible values are: open,
	// closed.  Default is "open".
	State string `url:"state,omitempty"`

	ListOptions
}

// Get list issues for gitlab projects
func (g *GitlabContext) ListIssues(project_id string, o *IssueListOptions) ([]*Issue, error) {
	path := getUrl([]string{"projects", url.QueryEscape(project_id), "issues"})
	u, err := addOptions(path, o)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("GET", u, nil)

	var ret []*Issue
	if _, err := g.Do(req, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

// Get list issues for gitlab projects
func (g *GitlabContext) CreateIssue(project_id string, issue *IssueRequest) (*Issue, int, error) {
	path := []string{"projects", url.QueryEscape(project_id), "issues"}
	req, _ := g.NewRequest("POST", path, issue)

	var ret *Issue
	if res, err := g.Do(req, &ret); err != nil {
		return nil, res.StatusCode, err
	}

	return ret, 0, nil
}

// Get list issues for gitlab projects
func (g *GitlabContext) UpdateIssue(project_id, issue_id string, issue *IssueRequest) (*Issue, int, error) {
	path := []string{"projects", url.QueryEscape(project_id), "issues", issue_id}
	req, _ := g.NewRequest("PUT", path, issue)

	var ret *Issue
	if res, err := g.Do(req, &ret); err != nil {
		return nil, res.StatusCode, err
	}

	return ret, 0, nil
}
