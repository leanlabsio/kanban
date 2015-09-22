package gitlab

import (
	"github.com/pmylund/sortutil"
	"net/http"
	"net/url"
	"time"
)

type Comment struct {
	Id        int64     `json:"id"`
	Author    *User     `json:"author"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

// ListComments returns list comments for gitlab issue
func (g *GitlabContext) ListComments(project_id, issue_id string, o *ListOptions) ([]*Comment, error) {
	cl := g.client
	path := getUrl([]string{"projects", url.QueryEscape(project_id), "issues", issue_id, "notes"})
	u, err := addOptions(path, o)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("GET", u, nil)

	var ret []*Comment
	if err := g.Do(cl, req, &ret); err != nil {
		return nil, err
	}

	sortutil.AscByField(ret, "CreatedAt")

	return ret, nil
}
