package gitlab

import (
	"net/http"
	"net/url"
)

// User represents a GitLab user.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/users.html
type User struct {
	Id           int64  `json:"id"`
	Name         string `json:"name,omitempty"`
	AvatarUrl    string `json:"avatar_url,nil,omitempty"`
	State        string `json:"state,omitempty"`
	Username     string `json:"username,omitempty"`
	WebUrl       string `json:"web_url,omitempty"`
	PrivateToken string `json:"private_token"`
}

// ListProjectMembers gets a list of a project's team members.
//
// GitLab API docs:
// http://doc.gitlab.com/ce/api/projects.html#list-project-team-members
func (g *GitlabContext) ListProjectMembers(project_id string, o *ListOptions) ([]*User, error) {
	path := getUrl([]string{"projects", url.QueryEscape(project_id), "members"})
	u, err := addOptions(path, o)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("GET", u, nil)

	var ret []*User
	if _, err := g.Do(req, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

// ListProjectMembers gets a list of a group's team members.
//
// GitLab API docs:
// http://doc.gitlab.com/ce/api/groups.html#list-group-members
func (g *GitlabContext) ListGroupMembers(group_id string, o *ListOptions) ([]*User, error)  {
	path := getUrl([]string{"groups", group_id, "members"})
	u, err := addOptions(path, o)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("GET", u, nil)

	var ret []*User
	if _, err := g.Do(req, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

// CurrentUser gets currently authenticated user.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/users.html#current-user
func (g *GitlabContext) CurrentUser() (*User, error) {
	path := getUrl([]string{"user"})
	req, _ := http.NewRequest("GET", path, nil)

	var ret *User
	if _, err := g.Do(req, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}
