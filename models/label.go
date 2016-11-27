package models

import (
	"regexp"
	"strconv"
)

// Label represent label
type Label struct {
	ID                    int64  `json:"id"`
	Color                 string `json:"color"`
	Name                  string `json:"name"`
	Description           string `json:"description"`
	OpenCardCount         int    `json:"open_card_count"`
	ClosedCardCount       int    `json:"closed_card_count"`
	OpenMergeRequestCount int    `json:"open_merge_requests_count"`
	Subscribed            bool   `json:"subscribed"`
	Priority              int    `json:"priority"`
}

// Stage represent board stage
type Stage struct {
	Name     string
	Position int
}

// LabelRequest represent request for update label
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
