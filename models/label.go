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

var (
	stageReg = regexp.MustCompile(`KB\[stage\]\[(\d)\]\[(.*)\]`)
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
