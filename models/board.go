package models
import (
	"gitlab.com/kanban/kanban/modules/gitlab"
)

type Board struct {
	Id                int64      `json:"id"`
	Name              string     `json:"name"`
	NamespaceWithName string     `json:"name_with_namespace"`
	PathWithNamespace string     `json:"path_with_namespace"`
	Namespace         *Namespace `json:"namespace,nil,omitempty"`
	Description       string     `json:"description"`
	LastModified      string     `json:"last_modified"`
	CreatedAt         string     `json:"created_at"`
	Owner             *User      `json:"owner,nil,omitempty"`
	AvatarUrl         string     `json:"avatar_url,nil,omitempty"`
}

type Namespace struct {
	Id     int64   `json:"id"`
	Name   string  `json:"name,omitempty"`
	Avatar *Avatar `json:"avatar,nil,omitempty"`
}

type Avatar struct {
	Url string `json:"url"`
}

func ListBoards(u *User, provider string) ([]*Board, error) {
	var b []*Board
	switch provider {
	case "gitlab":
		c := gitlab.NewContext(u.Credential["gitlab"].Token)
		r, err := c.ListProjects("100", "1")

		if err != nil {
			return nil, err
		}

		for _, item := range r {
			b = append(b, mapBoardFromGitlab(item))
		}
	}

	return b, nil
}

func ItemBoard(u *User, provider string, board_id string) (*Board, error) {
	var b *Board
	switch provider {
	case "gitlab":
		c := gitlab.NewContext(u.Credential["gitlab"].Token)
		r, err := c.ItemProject(board_id)

		if err != nil {
			return nil, err
		}
		b = mapBoardFromGitlab(r)
	}

	return b, nil
}

func mapBoardFromGitlab(r *gitlab.Project) *Board {
	return &Board{
		Id: r.Id,
		Name: r.Name,
		NamespaceWithName: r.NamespaceWithName,
		PathWithNamespace: r.PathWithNamespace,
		Namespace: mapNamespaceFromGitlab(r.Namespace),
		Description: r.Description,
		Owner: mapUserFromGitlab(r.Owner),
		AvatarUrl: r.AvatarUrl,
	}
}

func mapNamespaceFromGitlab(n *gitlab.Namespace) *Namespace {
	if n == nil {
		return nil
	}
	return &Namespace{
		Id: n.Id,
		Name: n.Name,
		Avatar: mapAvatarFromGitlab(n.Avatar),
	}
}

func mapAvatarFromGitlab(n *gitlab.Avatar) *Avatar {
	if n == nil {
		return nil
	}
	return &Avatar{
		Url: n.Url,
	}
}


