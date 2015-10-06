package gitlab

import (
	"net/http"
	"net/url"
)

// Issue represents a GitLab issue.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/issues.html
type Issue struct {
	Assignee    *User      `json:"assignee"`
	Author      *User      `json:"author"`
	Description string     `json:"description"`
	Milestone   *Milestone `json:"milestone"`
	Id          int64      `json:"id"`
	Iid         int64      `json:"iid"`
	Labels      *[]string   `json:"labels"`
	ProjectId   int64      `json:"project_id"`
	State       string     `json:"state"`
	Title       string     `json:"title"`
}

// IssueRequest represents the available CreateIssue() and UpdateIssue() options.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/issues.html#new-issues
type IssueRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	AssigneeId  *int64  `json:"assignee_id"`
	MilestoneId *int64  `json:"milestone_id"`
	Labels      string `json:"labels"`
	StateEvent  string `json:"state_event,omitempty"`
}

// ListIssuesOptions represents the available ListIssues() options.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/issues.html#list-issues
type IssueListOptions struct {
	// State filters issues based on their state.  Possible values are: open,
	// closed.  Default is "open".
	State string `url:"state,omitempty"`

	ListOptions
}

// ListIssues gets all issues created by authenticated user. This function
// takes pagination parameters page and per_page to restrict the list of issues.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/issues.html#list-issues
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

// CreateIssue creates a new project issue.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/issues.html#new-issues
func (g *GitlabContext) CreateIssue(project_id string, issue *IssueRequest) (*Issue, *http.Response, error) {
	path := []string{"projects", url.QueryEscape(project_id), "issues"}
	req, _ := g.NewRequest("POST", path, issue)

	var ret *Issue
	if res, err := g.Do(req, &ret); err != nil {
		return nil, res, err
	}

	return ret, nil, nil
}

// UpdateIssue updates an existing project issue. This function is also used
// to mark an issue as closed.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/issues.html#edit-issues
func (g *GitlabContext) UpdateIssue(project_id, issue_id string, issue *IssueRequest) (*Issue, *http.Response, error) {
	path := []string{"projects", url.QueryEscape(project_id), "issues", issue_id}
	req, _ := g.NewRequest("PUT", path, issue)

	var ret *Issue
	if res, err := g.Do(req, &ret); err != nil {
		return nil, res, err
	}

	return ret, nil, nil
}
