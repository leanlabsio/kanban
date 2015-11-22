package models

import (
	"gitlab.com/leanlabsio/kanban/modules/gitlab"
)

// Board represents a kanban board.
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

// Board represents a namespace kanban board.
type Namespace struct {
	Id     int64   `json:"id"`
	Name   string  `json:"name,omitempty"`
	Avatar *Avatar `json:"avatar,nil,omitempty"`
}

// Avatar represent a Avatar url.
type Avatar struct {
	Url string `json:"url"`
}

type BoardRequest struct {
	BoardId string `json:"project_id"`
}

var (
	defaultStages = []string{
		"KB[stage][0][Backlog]",
		"KB[stage][1][Development]",
		"KB[stage][2][Testing]",
		"KB[stage][3][Production]",
		"KB[stage][4][Ready]",
	}
)

// ListBoards returns list board for view user
func ListBoards(u *User, provider string) ([]*Board, error) {
	var b []*Board
	switch provider {
	case "gitlab":
		c := gitlab.NewContext(u.Credential["gitlab"].Token, u.Credential["gitlab"].PrivateToken)

		op := &gitlab.ProjectListOptions{}
		op.Page = "1"
		op.PerPage = "100"
		op.Archived = "false"

		r, err := c.ListProjects(op)

		if err != nil {
			return nil, err
		}

		for _, item := range r {
			b = append(b, mapBoardFromGitlab(item))
		}
	}

	return b, nil
}

// ItemBoard returns board item
func ItemBoard(u *User, provider string, board_id string) (*Board, error) {
	var b *Board
	switch provider {
	case "gitlab":
		c := gitlab.NewContext(u.Credential["gitlab"].Token, u.Credential["gitlab"].PrivateToken)
		r, err := c.ItemProject(board_id)

		if err, ok := err.(gitlab.ResponseError); ok {
			return nil, ReceivedDataErr{err.Error(), err.StatusCode}
		}

		if err != nil {
			return nil, err
		}
		b = mapBoardFromGitlab(r)
	}

	return b, nil
}

// ConfigureBoard creates default stages for board
func ConfigureBoard(u *User, provider string, f *BoardRequest) (int, error) {
	switch provider {
	case "gitlab":
		c := gitlab.NewContext(u.Credential["gitlab"].Token, u.Credential["gitlab"].PrivateToken)

		for _, stage := range defaultStages {
			_, res, err := c.CreateLabel(f.BoardId, &gitlab.LabelRequest{
				Name:  stage,
				Color: "#F5F5F5",
			})

			if err != nil {
				return res.StatusCode, err
			}
		}
	}

	return 0, nil
}

// mapBoardFromGitlab transforms board from gitlab project to kanban
func mapBoardFromGitlab(r *gitlab.Project) *Board {
	return &Board{
		Id:                r.Id,
		Name:              r.Name,
		NamespaceWithName: r.NamespaceWithName,
		PathWithNamespace: r.PathWithNamespace,
		Namespace:         mapNamespaceFromGitlab(r.Namespace),
		Description:       r.Description,
		Owner:             mapUserFromGitlab(r.Owner),
		AvatarUrl:         r.AvatarUrl,
	}
}

// mapNamespaceFromGitlab transforms namespace from gitlab to kanban
func mapNamespaceFromGitlab(n *gitlab.Namespace) *Namespace {
	if n == nil {
		return nil
	}
	return &Namespace{
		Id:     n.Id,
		Name:   n.Name,
		Avatar: mapAvatarFromGitlab(n.Avatar),
	}
}

// mapAvatarFromGitlab transform gitlab avatar to kanban avatar
func mapAvatarFromGitlab(n *gitlab.Avatar) *Avatar {
	if n == nil {
		return nil
	}
	return &Avatar{
		Url: n.Url,
	}
}
