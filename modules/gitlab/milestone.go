package gitlab

import _ "encoding/json"

type Milestone struct {
	Id    int64  `json:"id"`
	State string `json:"state,omitempty"`
	Title string `json:"title,omitempty"`
}

type MilestoneListResponse struct {
	Data []Milestone `json:"data"`
	Meta []string    `json:"meta"`
}
