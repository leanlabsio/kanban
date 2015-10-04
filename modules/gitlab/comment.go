package gitlab

import (
	"github.com/pmylund/sortutil"
	"net/http"
	"net/url"
	"time"
)

// Comment represents a GitLab note.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/notes.html
type Comment struct {
	Id        int64     `json:"id"`
	Author    *User     `json:"author"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

// CommentRequest represents the available CreateComment() and UpdateComment()
// options.
//
// GitLab API docs:
// http://doc.gitlab.com/ce/api/notes.html#create-new-issue-note
type CommentRequest struct {
	Body string `json:"body"`
}

// ListComments gets a list of all notes for a single issue.
//
// GitLab API docs:
// http://doc.gitlab.com/ce/api/notes.html#list-project-issue-notes
func (g *GitlabContext) ListComments(project_id, issue_id string, o *ListOptions) ([]*Comment, error) {
	path := getUrl([]string{"projects", url.QueryEscape(project_id), "issues", issue_id, "notes"})
	u, err := addOptions(path, o)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("GET", u, nil)

	var ret []*Comment
	if _, err := g.Do(req, &ret); err != nil {
		return nil, err
	}

	sortutil.Reverse(ret)

	return ret, nil
}

// CreateComment creates a new note to a single project issue.
//
// GitLab API docs:
// http://doc.gitlab.com/ce/api/notes.html#create-new-issue-note
func (g *GitlabContext) CreateComment(project_id, issue_id string, com *CommentRequest) (*Comment, *http.Response, error) {
	path := []string{"projects", url.QueryEscape(project_id), "issues", issue_id, "notes"}
	req, _ := g.NewRequest("POST", path, com)

	var ret *Comment
	if res, err := g.Do(req, &ret); err != nil {
		return nil, res, err
	}

	return ret, nil, nil
}
