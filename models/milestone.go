package models

// Milestone represents a kanban milestone
type Milestone struct {
	ID          int64  `json:"id"`
	IID         int64  `json:"iid"`
	State       string `json:"state,omitempty"`
	Title       string `json:"title,omitempty"`
	DueDate     string `json:"due_date,omitempty"`
	Description string `json:"description"`
	UpdatedAt   string `json:"updated_at"`
	CreatedAt   string `json:"created_at"`
}

// MilestoneRequest represents a milestone request for create, update, delete milestone on kanban
type MilestoneRequest struct {
	MilestoneID int64  `json:"milestone_id"`
	ProjectID   int64  `json:"project_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
}
