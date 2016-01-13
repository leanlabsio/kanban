package models

import (
	"encoding/json"
	"time"
)

// Comment represents a card comment
type Comment struct {
	Id        int64     `json:"id"`
	Author    *User     `json:"author"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	IsInfo    bool      `json:"is_info"`
}

// CommentRequest represents a request for create or update comment on card
type CommentRequest struct {
	CardId    int64  `json:"issue_id"`
	ProjectId int64  `json:"project_id"`
	Body      string `json:"body"`
}

// Marshal returns the JSON encoding of comment
func (c *Comment) MarshalJSON() ([]byte, error) {
	type Alias Comment
	return json.Marshal(struct {
		CreatedAt int64 `json:"created_at"`
		*Alias
	}{
		CreatedAt: c.CreatedAt.Unix(),
		Alias:     (*Alias)(c),
	})
}
