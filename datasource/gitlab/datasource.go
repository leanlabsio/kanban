package gitlab

import (
	"golang.org/x/oauth2"
	gitlabclient "gitlab.com/leanlabsio/kanban/modules/gitlab" 
)

type GitLabDataSource struct {
	client *gitlabclient.GitlabContext 
}

func New(t *oauth2.Token, pt string) GitLabDataSource {
	c := gitlabclient.NewContext(t, pt)

	return GitLabDataSource{client: c}
}
