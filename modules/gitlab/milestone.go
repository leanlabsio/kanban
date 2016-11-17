package gitlab

import (
	"net/http"
	"net/url"
)

// Milestone represents a GitLab milestone.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/branches.html
type Milestone struct {
	ID          int64  `json:"id"`
	IID         int64  `json:"iid"`
	State       string `json:"state,omitempty"`
	Title       string `json:"title,omitempty"`
	DueDate     string `json:"due_date,omitempty"`
	Description string `json:"description"`
	UpdatedAt   string `json:"updated_at"`
	CreatedAt   string `json:"created_at"`
}

// MilestoneRequest represents the available CreateMilestone() and UpdateMilestone() options.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/milestones.html#create-new-milestone
type MilestoneRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
}

// ListMilestones returns a list of project milestones.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/milestones.html#list-project-milestones
func (g *GitlabContext) ListMilestones(project_id string, o *ListOptions) ([]*Milestone, error) {
	path := getUrl([]string{"projects", url.QueryEscape(project_id), "milestones"})
	u, err := addOptions(path, o)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("GET", u, nil)

	var ret []*Milestone
	if _, err := g.Do(req, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

// CreateMilestone creates a new project milestone.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/milestones.html#create-new-milestone
func (g *GitlabContext) CreateMilestone(project_id string, m *MilestoneRequest) (*Milestone, *http.Response, error) {
	path := []string{"projects", url.QueryEscape(project_id), "milestones"}
	req, _ := g.NewRequest("POST", path, m)

	var ret *Milestone
	if res, err := g.Do(req, &ret); err != nil {
		return nil, res, err
	}

	return ret, nil, nil
}

// UpdateMilestone updates a project milestone.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/milestones.html#edit-milestone
func (g *GitlabContext) UpdateMilestone(project_id string, m_id string, m *MilestoneRequest) (*Milestone, *http.Response, error) {
	path := []string{"projects", url.QueryEscape(project_id), "milestones", m_id}
	req, _ := g.NewRequest("PUT", path, m)

	var ret *Milestone
	if res, err := g.Do(req, &ret); err != nil {
		return nil, res, err
	}

	return ret, nil, nil
}
