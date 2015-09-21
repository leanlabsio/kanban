package models

import (
	"time"
	"gitlab.com/kanban/kanban/modules/gitlab"
	"encoding/json"
	"regexp"
	"strings"
)

type Comment struct {
	Id        int64     `json:"id"`
	Author    *User     `json:"author"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	IsInfo    bool      `json:"is_info"`
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
	}, "|"))

	regInfo = regexp.MustCompile("^" + reg + "$")
)

// ListComments returns list comments for card
func ListComments(u *User, provider, project_id, card_id string) ([]*Comment, error) {
	var b []*Comment
	switch provider {
	case "gitlab":
		c := gitlab.NewContext(u.Credential["gitlab"].Token)
		r, err := c.ListComments(project_id, card_id)

		if err != nil {
			return nil, err
		}

		b = mapCommentCollectionFromGitlab(r)
	}

	return b, nil
}

// mapCommentCollectionFromGitlab transforms gitlab coments to kanban comments
func mapCommentCollectionFromGitlab(c []*gitlab.Comment) []*Comment {
	var b []*Comment
	for _, co := range (c) {
		b = append(b, mapCommentFromGitlab(co))
	}

	return b
}

// mapCommentFromGitlab transform gitlab comment to kanban comment
func mapCommentFromGitlab(c *gitlab.Comment) *Comment {
	return &Comment{
		Id: c.Id,
		Author: mapUserFromGitlab(c.Author),
		Body: c.Body,
		CreatedAt: c.CreatedAt,
		IsInfo: mapCommentIsInfoFromGitlab(c.Body),
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