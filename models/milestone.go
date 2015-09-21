package models

import (
	"gitlab.com/kanban/kanban/modules/gitlab"
)

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
		c := gitlab.NewContext(u.Credential["gitlab"].Token)
		r, err := c.ListMilestones(board_id)

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
	return &Milestone{
		Id:    m.Id,
		State: m.State,
		Title: m.Title,
	}
}
