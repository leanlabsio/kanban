package gitlab

type Issue struct {
	Assignee    *User       `json:"assignee"`
	Author      *User       `json:"author"`
	Description string      `json:"description"`
	Id          int64       `json:"id"`
	Iid         int64       `json:"iid"`
	Labels      []string    `json:"labels"`
	ProjectId   int64       `json:"project_id"`
	State       string      `json:"state"`
	Title       string      `json:"title"`
}
