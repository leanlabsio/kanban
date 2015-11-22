package models

import (
	"gitlab.com/leanlabsio/kanban/modules/gitlab"
)

// Milestone represents a kanban milestone
type Milestone struct {
	Id    int64  `json:"id"`
	State string `json:"state,omitempty"`
	Title string `json:"title,omitempty"`
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

// mapMilestoneFromGitlab returns map from gitlab milestone to gitlab milestone
func mapMilestoneFromGitlab(m *gitlab.Milestone) *Milestone {
	if m == nil {
		return nil
	}

	return &Milestone{
		Id:    m.Id,
		State: m.State,
		Title: m.Title,
	}
}
