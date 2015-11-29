package models

import (
	"gitlab.com/leanlabsio/kanban/modules/gitlab"
	"fmt"
)

// Milestone represents a kanban milestone
type Milestone struct {
	ID    int64    `json:"id"`
	State string   `json:"state,omitempty"`
	Title string   `json:"title,omitempty"`
	DueDate string `json:"due_date,omitempty"`
}


// MilestoneRequest represents a milestone request for create, update, delete milestone on kanban
type MilestoneRequest struct {
	MilestoneID int64  `json:"milestone_id"`
	ProjectID   int64  `json:"project_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
}

// ListMilestones returns list milestones by project
func ListMilestones(u *User, provider, board_id string) ([]*Milestone, error) {
	var mem []*Milestone
	switch provider {
	case "gitlab":
		c := gitlab.NewContext(u.Credential["gitlab"].Token, u.Credential["gitlab"].PrivateToken)
		r, err := c.ListMilestones(board_id, &gitlab.ListOptions{
			Page:    "1",
			PerPage: "100",
		})

		if err != nil {
			return nil, err
		}

		for _, item := range r {
			mem = append(mem, mapMilestoneFromGitlab(item))
		}
	}

	return mem, nil
}

// CreateMilestone create new milestone on board
func CreateMilestone(u *User, provider string, form *MilestoneRequest) (*Milestone, int, error)  {
	var cr *Milestone
	var code int
	switch provider {
	case "gitlab":
		c := gitlab.NewContext(u.Credential["gitlab"].Token, u.Credential["gitlab"].PrivateToken)
		r, res, err := c.CreateMilestone(fmt.Sprintf("%d", form.ProjectID), mapMilestoneRequestToGitlab(form))
		if err != nil {
			return nil, res.StatusCode, err
		}

		cr = mapMilestoneFromGitlab(r)
	}

	return cr, code, nil
}

// mapMilestoneRequestToGitlab transforms kanban milestone to gitlab milestone request
func mapMilestoneRequestToGitlab(m *MilestoneRequest) *gitlab.MilestoneRequest {
	return &gitlab.MilestoneRequest{
		Title:       m.Title,
		Description: m.Description,
		DueDate:     m.DueDate,
	}
}

// mapMilestoneFromGitlab returns map from gitlab milestone to gitlab milestone
func mapMilestoneFromGitlab(m *gitlab.Milestone) *Milestone {
	if m == nil {
		return nil
	}

	return &Milestone{
		ID:      m.ID,
		State:   m.State,
		Title:   m.Title,
		DueDate: m.DueDate,
	}
}
