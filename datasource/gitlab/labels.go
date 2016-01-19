package gitlab

import (
	"gitlab.com/leanlabsio/kanban/models"
	"gitlab.com/leanlabsio/kanban/modules/gitlab"
)

// ListLabels returns list kanban labels for board
func (ds GitLabDataSource) ListLabels(board_id string) ([]*models.Label, error) {
	var l []*models.Label
	r, err := ds.client.ListLabels(board_id, &gitlab.ListOptions{
		Page:    "1",
		PerPage: "100",
	})

	if err != nil {
		return nil, err
	}

	for _, v := range r {
		l = append(l, mapLabelFromGitlab(v))
	}

	return l, nil
}

// mapLabelFromGitlab transforms gitlab label to kanban label
func mapLabelFromGitlab(l *gitlab.Label) *models.Label {
	return &models.Label{
		Color: l.Color,
		Name:  l.Name,
	}
}
