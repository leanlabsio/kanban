package gitlab

import (
	"gitlab.com/leanlabsio/kanban/modules/gitlab"
	"gitlab.com/leanlabsio/kanban/models"
	"fmt"
)

// MilestoneRequest represents a milestone request for create, update, delete milestone on kanban
type MilestoneRequest struct {
	MilestoneID int64  `json:"milestone_id"`
	ProjectID   int64  `json:"project_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
}

// ListMilestones returns list milestones by project
func (ds GitLabDataSource) ListMilestones(board_id string) ([]*models.Milestone, error) {
	var mem []*models.Milestone
	r, err := ds.client.ListMilestones(board_id, &gitlab.ListOptions{
		Page:    "1",
		PerPage: "100",
	})

	if err != nil {
		return nil, err
	}

	for _, item := range r {
		mem = append(mem, mapMilestoneFromGitlab(item))
	}

	return mem, nil
}

// CreateMilestone create new milestone on board
func (ds GitLabDataSource) CreateMilestone(form *MilestoneRequest) (*models.Milestone, int, error)  {
	var cr *models.Milestone
	var code int
	r, res, err := ds.client.CreateMilestone(fmt.Sprintf("%d", form.ProjectID), mapMilestoneRequestToGitlab(form))
	if err != nil {
		return nil, res.StatusCode, err
	}

	cr = mapMilestoneFromGitlab(r)

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
func mapMilestoneFromGitlab(m *gitlab.Milestone) *models.Milestone {
	if m == nil {
		return nil
	}

	return &models.Milestone{
		ID:      m.ID,
		State:   m.State,
		Title:   m.Title,
		DueDate: m.DueDate,
	}
}
