package gitlab

import (
	gitlabclient "gitlab.com/leanlabsio/kanban/modules/gitlab"
	"golang.org/x/oauth2"
	"gopkg.in/redis.v3"
)

type GitLabDataSource struct {
	client *gitlabclient.GitlabContext
	db     *redis.Client
}

// New create new gitlab datasource instance
func New(t *oauth2.Token, pt string, r *redis.Client) GitLabDataSource {
	c := gitlabclient.NewContext(t, pt)

	return GitLabDataSource{client: c, db: r}
}
