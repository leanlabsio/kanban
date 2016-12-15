package gitlab

import (
	"regexp"
	"strconv"
	"strings"

	"gitlab.com/leanlabsio/kanban/models"
	"gitlab.com/leanlabsio/kanban/modules/gitlab"
)

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
		`Marked the task .*? as completed`,
		`Marked the task .*? as incomplete`,
	}, "|"))

	regInfo = regexp.MustCompile("^" + reg + "$")
)

// ListComments gets a list of all comment for a single card.
func (ds GitLabDataSource) ListComments(project_id, card_id string) ([]*models.Comment, error) {
	b := make([]*models.Comment, 0)
	r, err := ds.client.ListComments(project_id, card_id, &gitlab.ListOptions{
		Page:    "1",
		PerPage: "100",
	})

	if err != nil {
		return nil, err
	}

	for _, co := range r {
		b = append(b, mapCommentFromGitlab(co))
	}

	return b, nil
}

// CreateComment creates a new comment to a single board card.
func (ds GitLabDataSource) CreateComment(form *models.CommentRequest) (*models.Comment, int, error) {
	var b *models.Comment
	var code int
	r, res, err := ds.client.CreateComment(
		strconv.FormatInt(form.ProjectId, 10),
		strconv.FormatInt(form.CardId, 10),
		mapCommentRequestToGitlab(form),
	)

	if err != nil {
		return nil, res.StatusCode, err
	}

	b = mapCommentFromGitlab(r)

	return b, code, nil
}

// mapCommentRequestToGitlab transforms kanban comment request to gitlab comment request
func mapCommentRequestToGitlab(c *models.CommentRequest) *gitlab.CommentRequest {
	return &gitlab.CommentRequest{
		Body: c.Body,
	}
}

// mapCommentFromGitlab transform gitlab comment to kanban comment
func mapCommentFromGitlab(c *gitlab.Comment) *models.Comment {
	return &models.Comment{
		Id:        c.Id,
		Author:    mapUserFromGitlab(c.Author),
		Body:      c.Body,
		CreatedAt: c.CreatedAt,
		IsInfo:    mapCommentIsInfoFromGitlab(c),
	}
}

// mapCommentIsInfoFromGitlab checks type comment from gitlab comment body
func mapCommentIsInfoFromGitlab(c *gitlab.Comment) bool {
	return regInfo.MatchString(c.Body) || c.System
}
