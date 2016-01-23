package gitlab

import (
	"fmt"
	"gitlab.com/leanlabsio/kanban/models"
	"gitlab.com/leanlabsio/kanban/modules/gitlab"
)

// ListMembers is
func (ds GitLabDataSource) ListMembers(board_id string) ([]*models.User, error) {
	var mem []*models.User
	r, err := ds.client.ListProjectMembers(board_id, &gitlab.ListOptions{
		Page:    "1",
		PerPage: "100",
	})

	if err != nil {
		return nil, err
	}

	b, err := ds.ItemBoard(board_id)
	fmt.Printf("%+v", b)

	if err != nil {
		return nil, err
	}

	exist := make(map[string]bool)

	if b.Owner == nil {
		u, _ := ds.client.ListGroupMembers(fmt.Sprintf("%d", b.Namespace.Id), &gitlab.ListOptions{})

		for _, item := range u {
			exist[item.Username] = true
			mem = append(mem, mapUserFromGitlab(item))
		}
	}

	for _, item := range r {
		if !exist[item.Username] {
			mem = append(mem, mapUserFromGitlab(item))
		}
	}

	return mem, nil
}

// mapUserFromGitlab mapped data from gitlab user to kanban user
func mapUserFromGitlab(u *gitlab.User) *models.User {
	if u == nil {
		return nil
	}
	return &models.User{
		Id:        u.Id,
		Name:      u.Name,
		Username:  u.Username,
		AvatarUrl: u.AvatarUrl,
		State:     u.State,
	}
}
