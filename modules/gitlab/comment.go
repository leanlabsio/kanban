package gitlab

import (
	"net/http"
	"net/url"
	"time"
	"sort"
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

// commentSlice represents list comments for usage sort.Interface
type commentSlice []*Comment

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

	var ret commentSlice
	if _, err := g.Do(req, &ret); err != nil {
		return nil, err
	}

	sort.Sort(ret)

	return ret, nil
}

// Len is the number of elements in the collection.
func (p commentSlice) Len() int {
	return len(p)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (p commentSlice) Less(i, j int) bool {
	return p[i].CreatedAt.Before(p[j].CreatedAt)
}

// Swap swaps the elements with indexes i and j.
func (p commentSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
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
