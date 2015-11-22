package models

import (
	"gitlab.com/leanlabsio/kanban/modules/gitlab"
	"regexp"
	"strconv"
)

type Label struct {
	Color string `json:"color"`
	Name  string `json:"name"`
}

type Stage struct {
	Name     string
	Position int
}

var (
	stageReg = regexp.MustCompile(`KB\[stage\]\[(\d)\]\[(.*)\]`)
)

// ListLabels returns list kanban labels for board
func ListLabels(u *User, provider, board_id string) ([]*Label, error) {
	var l []*Label
	switch provider {
	case "gitlab":
		c := gitlab.NewContext(u.Credential["gitlab"].Token, u.Credential["gitlab"].PrivateToken)
		r, err := c.ListLabels(board_id, &gitlab.ListOptions{
			Page:    "1",
			PerPage: "100",
		})

		if err != nil {
			return nil, err
		}

		for _, v := range r {
			l = append(l, mapLabelFromGitlab(v))
		}
	}

	return l, nil
}

// ParseLabelToStage transform label to stage
func ParseLabelToStage(l string) *Stage {
	m := stageReg.MatchString(l)

	var s Stage
	if m {
		an := stageReg.FindStringSubmatch(l)
		s.Position, _ = strconv.Atoi(an[1])
		s.Name = an[2]
	}

	return &s
}

// mapLabelFromGitlab transforms gitlab label to kanban label
func mapLabelFromGitlab(l *gitlab.Label) *Label {
	return &Label{
		Color: l.Color,
		Name:  l.Name,
	}
}
