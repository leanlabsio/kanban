package models

import (
	"fmt"
)

// Card represents an card in kanban board
type Card struct {
	Id          int64       `json:"id"`
	Iid         int64       `json:"iid"`
	Assignee    *User       `json:"assignee"`
	Milestone   *Milestone  `json:"milestone"`
	Author      *User       `json:"author"`
	Description string      `json:"description"`
	Labels      *[]string   `json:"labels"`
	ProjectId   int64       `json:"project_id"`
	Properties  *Properties `json:"properties"`
	State       string      `json:"state"`
	Title       string      `json:"title"`
	Todo        []*Todo     `json:"todo"`
}

// Properties represents a card properties
type Properties struct {
	Andon string `json:"andon"`
}

// Todo represents an todo an card
type Todo struct {
	Body    string `json:"body"`
	Checked bool   `json:"checked"`
}

// CardRequest represents a card request for create, update, delete card on kanban
type CardRequest struct {
	CardId      int64             `json:"issue_id"`
	ProjectId   int64             `json:"project_id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	AssigneeId  *int64            `json:"assignee_id"`
	MilestoneId *int64            `json:"milestone_id"`
	Labels      string            `json:"labels"`
	Properties  *Properties       `json:"properties"`
	Stage       map[string]string `json:"stage"`
	Todo        []*Todo           `json:"todo"`
}

func (c *Card) RoutingKey() string {
	return fmt.Sprintf("kanban.%d", c.ProjectId)
}
