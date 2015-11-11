package gitlab

import (
	"net/http"
	"net/url"
	"strings"
)

// Project represents a GitLab project.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/projects.html
type Project struct {
	Id                int64      `json:"id"`
	Name              string     `json:"name"`
	NamespaceWithName string     `json:"name_with_namespace"`
	PathWithNamespace string     `json:"path_with_namespace"`
	Namespace         *Namespace `json:"namespace,nil,omitempty"`
	Description       string     `json:"description"`
	LastModified      string     `json:"last_modified"`
	CreatedAt         string     `json:"created_at"`
	Owner             *User      `json:"owner,nil,omitempty"`
	AvatarUrl         string     `json:"avatar_url,nil,omitempty"`
}

// Namespace represents a GitLab namespace.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/namespaces.html
type Namespace struct {
	Id     int64   `json:"id"`
	Name   string  `json:"name,omitempty"`
	Avatar *Avatar `json:"avatar,nil,omitempty"`
}

// Avatar represents a GitLab avatar.
type Avatar struct {
	Url string `json:"url"`
}

// ProjectListOptions represents the available ListProjects() options.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/projects.html#list-projects
type ProjectListOptions struct {
	// State filters issues based on their state.  Possible values are: open,
	// closed.  Default is "open".
	Archived string `url:"archived,omitempty"`

	ListOptions
}

// ListProjects gets a list of projects accessible by the authenticated user.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/projects.html#list-projects
func (g *GitlabContext) ListProjects(o *ProjectListOptions) ([]*Project, error) {
	u, err := addOptions(getUrl([]string{"projects"}), o)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("GET", u, nil)

	var ret []*Project
	if _, err := g.Do(req, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

// ItemProject gets a specific project, identified by project ID or
// NAMESPACE/PROJECT_NAME, which is owned by the authenticated user.
//
// GitLab API docs:
// http://doc.gitlab.com/ce/api/projects.html#get-single-project
func (g *GitlabContext) ItemProject(project_id string) (*Project, error) {
	path := getUrl([]string{"projects", strings.Replace(url.QueryEscape(project_id), ".", "%2E", -1)})
	req, _ := http.NewRequest("GET", path, nil)

	var ret Project
	if _, err := g.Do(req, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
