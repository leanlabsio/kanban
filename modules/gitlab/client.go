package gitlab

import (
	"github.com/pmylund/sortutil"
	"golang.org/x/oauth2"
	"net/http"
	"net/url"
)

type GitlabClient struct {
}

type Config struct {
	BasePath string
	Domain   string
	Oauth2   *oauth2.Config
}

type GitlabContext struct {
	client *http.Client
}

var (
	cfg *Config
)

// New gitlab api client
func NewEngine(c *Config) {
	cfg = c
}

// AuthCodeURL returns a URL to OAuth 2.0 provider's consent page
// that asks for permissions for the required scopes explicitly.
func AuthCodeURL() string {
	return  cfg.Oauth2.AuthCodeURL("state", oauth2.AccessTypeOffline)
}

// Exchange is
func Exchange(c string) (*oauth2.Token, error) {
	return cfg.Oauth2.Exchange(oauth2.NoContext, c)
}

// NewContext
func NewContext(t *oauth2.Token) (*GitlabContext) {
	return &GitlabContext{
		client: cfg.Oauth2.Client(oauth2.NoContext, t),
	}
}

// List projects from gitlab
func (g *GitlabContext) ListProjects(per_page, page string) ([]*Project, error) {
	cl := g.client
	path := g.GetUrl([]string{"projects"})

	req, _ := http.NewRequest("GET", path, nil)
	q := req.URL.Query()
	q.Add("per_page", per_page)
	q.Add("page", page)
	q.Add("archived", "false")
	req.URL.RawQuery = q.Encode()

	var ret []*Project
	if err := g.Do(cl, req, &ret); err != nil {
		return  nil, err
	}

	return ret, nil
}

// ItemProject returns project item from gitlab
func (g *GitlabContext) ItemProject(project_id string) (*Project, error) {
	cl := g.client
	path := g.GetUrl([]string{"projects", url.QueryEscape(project_id)})

	req, _ := http.NewRequest("GET", path, nil)

	var ret Project
	if err := g.Do(cl, req, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

// Get list issues for gitlab projects
func (g *GitlabContext) ListIssues(project_id, per_page, page string) ([]*Issue, error) {
	cl := g.client
	path := g.GetUrl([]string{"projects", url.QueryEscape(project_id), "issues"})

	req, _ := http.NewRequest("GET", path, nil)
	q := req.URL.Query()
	q.Add("per_page", per_page)
	q.Add("page", page)
	req.URL.RawQuery = q.Encode()

	var ret []*Issue
	if err := g.Do(cl, req, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

// Get list milestones for gitlab projects
func (g *GitlabContext) ListMilestones(project_id string) ([]*Milestone, error) {
	cl := g.client
	path := g.GetUrl([]string{"projects", url.QueryEscape(project_id), "milestones"})

	req, _ := http.NewRequest("GET", path, nil)
	q := req.URL.Query()
	q.Add("per_page", "100")
	q.Add("page", "1")
	req.URL.RawQuery = q.Encode()

	var ret []*Milestone
	if err := g.Do(cl, req, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

// Get list project members for gitlab projects
func (g *GitlabContext) ListProjectMembers(project_id string) ([]*User, error) {
	cl := g.client
	path := g.GetUrl([]string{"projects", url.QueryEscape(project_id), "members"})

	req, _ := http.NewRequest("GET", path, nil)
	q := req.URL.Query()
	q.Add("per_page", "100")
	q.Add("page", "1")
	req.URL.RawQuery = q.Encode()

	var ret []*User
	if err := g.Do(cl, req, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

// ListLabels returns list labels for gitlab projects
func (g *GitlabContext) ListLabels(project_id string) ([]*Label, error) {
	cl := g.client
	path := g.GetUrl([]string{"projects", url.QueryEscape(project_id), "labels"})

	req, _ := http.NewRequest("GET", path, nil)

	var ret []*Label
	if err := g.Do(cl, req, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

// ListComments returns list comments for gitlab issue
func (g *GitlabContext) ListComments(project_id, issue_id string) ([]*Comment, error) {
	cl := g.client
	path := g.GetUrl([]string{"projects", url.QueryEscape(project_id), "issues", issue_id, "notes"})

	req, _ := http.NewRequest("GET", path, nil)

	var ret []*Comment
	if err := g.Do(cl, req, &ret); err != nil {
		return nil, err
	}

	sortutil.AscByField(ret, "CreatedAt")

	return ret, nil
}

// CurrentUser returns current authentificated user
func (g *GitlabContext) CurrentUser() (*User, error) {
	cl := g.client
	path := g.GetUrl([]string{"user"})

	req, _ := http.NewRequest("GET", path, nil)

	var ret *User
	if err := g.Do(cl, req, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}
