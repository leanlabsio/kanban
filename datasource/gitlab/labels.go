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

// EditLabel updates existing label
func (ds GitLabDataSource) EditLabel(project_id string, req *models.LabelRequest) (*models.Label, error) {
	var l *models.Label
	r, _, err := ds.client.EditLabel(project_id, &gitlab.LabelRequest{
		Name:    req.Name,
		Color:   req.Color,
		NewName: req.NewName,
	})

	if err != nil {
		return nil, err
	}
	l = mapLabelFromGitlab(r)

	return l, nil
}

func (ds GitLabDataSource) DeleteLabel(project_id, name string) (*models.Label, error) {
	var l *models.Label

	r, _, err := ds.client.DeleteLabel(project_id, &gitlab.LabelRequest{
		Name: name,
	})

	if err != nil {
		return nil, err
	}
	l = mapLabelFromGitlab(r)

	return l, nil
}

func (ds GitLabDataSource) CreateLabel(project_id string, req *models.LabelRequest) (*models.Label, error) {
	var l *models.Label
	r, _, err := ds.client.CreateLabel(project_id, &gitlab.LabelRequest{
		Name:  req.Name,
		Color: req.Color,
	})
	if err != nil {
		return nil, err
	}
	l = mapLabelFromGitlab(r)

	return l, nil
}

// mapLabelFromGitlab transforms gitlab label to kanban label
func mapLabelFromGitlab(l *gitlab.Label) *models.Label {
	return &models.Label{
		Color: l.Color,
		Name:  l.Name,
	}
}
