package gitlab

import (
	gitlabclient "gitlab.com/leanlabsio/kanban/modules/gitlab"
	"golang.org/x/oauth2"
)

type GitLabDataSource struct {
	client *gitlabclient.GitlabContext
}

func New(t *oauth2.Token, pt string) GitLabDataSource {
	c := gitlabclient.NewContext(t, pt)

	return GitLabDataSource{client: c}
}
