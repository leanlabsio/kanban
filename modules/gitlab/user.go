package gitlab

import (
	"net/http"
	"net/url"
)

type User struct {
	Id        int64  `json:"id"`
	Name      string `json:"name,omitempty"`
	AvatarUrl string `json:"avatar_url,nil,omitempty"`
	State     string `json:"state,omitempty"`
	Username  string `json:"username,omitempty"`
	WebUrl    string `json:"web_url,omitempty"`
}

// Get list project members for gitlab projects
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

// CurrentUser returns current authentificated user
func (g *GitlabContext) CurrentUser() (*User, error) {
	path := getUrl([]string{"user"})
	req, _ := http.NewRequest("GET", path, nil)

	var ret *User
	if _, err := g.Do(req, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}
