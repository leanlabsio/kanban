package models

import (
	"encoding/json"
	"gitlab.com/leanlabsio/kanban/modules/gitlab"
	"regexp"
	"strconv"
	"strings"
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

var (
	reg = string(strings.Join([]string{
		`Reassigned to .*?`,
		`Milestone changed to .*?`,
		`Title changed from .*? to .*?`,
		`Added .*? label(s)?`,
		`mentioned in commit .*?`,
		`mentioned in merge request .*?`,
		`Status changed to closed`,
		`Status changed to reopened`,
		`moved issue from .*? to .*?`,
		`Marked as \*\*blocked\*\*: .*?`,
		`Marked as \*\*ready\*\* for next stage`,
		`Marked as \*\*unblocked\*\*`,
		`mentioned in issue .*?`,
		`Assignee removed`,
		`Milestone removed`,
	}, "|"))

	regInfo = regexp.MustCompile("^" + reg + "$")
)

// ListComments gets a list of all comment for a single card.
func ListComments(u *User, provider, project_id, card_id string) ([]*Comment, error) {
	b := make([]*Comment, 0)
	switch provider {
	case "gitlab":
		c := gitlab.NewContext(u.Credential["gitlab"].Token, u.Credential["gitlab"].PrivateToken)
		r, err := c.ListComments(project_id, card_id, &gitlab.ListOptions{
			Page:    "1",
			PerPage: "100",
		})

		if err != nil {
			return nil, err
		}

		for _, co := range r {
			b = append(b, mapCommentFromGitlab(co))
		}
	}

	return b, nil
}

// CreateComment creates a new comment to a single board card.
func CreateComment(u *User, provider string, form *CommentRequest) (*Comment, int, error) {
	var b *Comment
	var code int
	switch provider {
	case "gitlab":
		c := gitlab.NewContext(u.Credential["gitlab"].Token, u.Credential["gitlab"].PrivateToken)
		r, res, err := c.CreateComment(
			strconv.FormatInt(form.ProjectId, 10),
			strconv.FormatInt(form.CardId, 10),
			mapCommentRequestToGitlab(form),
		)

		if err != nil {
			return nil, res.StatusCode, err
		}

		b = mapCommentFromGitlab(r)
	}

	return b, code, nil
}

// mapCommentRequestToGitlab transforms kanban comment request to gitlab comment request
func mapCommentRequestToGitlab(c *CommentRequest) *gitlab.CommentRequest {
	return &gitlab.CommentRequest{
		Body: c.Body,
	}
}

// mapCommentFromGitlab transform gitlab comment to kanban comment
func mapCommentFromGitlab(c *gitlab.Comment) *Comment {
	return &Comment{
		Id:        c.Id,
		Author:    mapUserFromGitlab(c.Author),
		Body:      c.Body,
		CreatedAt: c.CreatedAt,
		IsInfo:    mapCommentIsInfoFromGitlab(c.Body),
	}
}

// mapCommentIsInfoFromGitlab checks type comment from gitlab comment body
func mapCommentIsInfoFromGitlab(b string) bool {
	return regInfo.MatchString(b)
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
