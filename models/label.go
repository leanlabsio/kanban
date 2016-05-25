package models

import (
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

type LabelRequest struct {
	Name    string `json:"name"`
	Color   string `json:"color"`
	NewName string `json:"new_name"`
}

var (
	stageReg = regexp.MustCompile(`KB\[stage\]\[(\d+)\]\[(.*)\]`)
)

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
