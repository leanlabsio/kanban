package models

import "gitlab.com/kanban/kanban/modules/gitlab"

type Label struct {
	Color string `json:"color"`
	Name  string `json:"name"`
}

// ListLabels returns list kanban labels for board
func ListLabels(u *User, provider, board_id string) ([]*Label, error) {
	var l []*Label
	switch provider {
	case "gitlab":
		c := gitlab.NewContext(u.Credential["gitlab"].Token)
		r, err := c.ListLabels(board_id, &gitlab.ListOptions{
			Page:    "1",
			PerPage: "100",
		})

		if err != nil {
			return nil, err
		}

		l = mapLabelCollectionFromGitlab(r)
	}

	return l, nil
}

// mapLabelCollectionFromGitlab transforms gitlab labels to kanban labels
func mapLabelCollectionFromGitlab(l []*gitlab.Label) []*Label {
	var b []*Label
	for _, v := range l {
		b = append(b, mapLabelFromGitlab(v))
	}

	return b
}

// mapLabelFromGitlab transforms gitlab label to kanban label
func mapLabelFromGitlab(l *gitlab.Label) *Label {
	return &Label{
		Color: l.Color,
		Name:  l.Name,
	}
}
