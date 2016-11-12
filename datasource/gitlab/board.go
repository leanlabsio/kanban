package gitlab

import (
	"fmt"

	"gitlab.com/leanlabsio/kanban/models"
	"gitlab.com/leanlabsio/kanban/modules/gitlab"
)

var (
	defaultStages = []string{
		"KB[stage][10][Backlog]",
		"KB[stage][20][Development]",
		"KB[stage][30][Testing]",
		"KB[stage][40][Production]",
		"KB[stage][50][Ready]",
	}
)

// ListBoards returns list board for view user
func (ds GitLabDataSource) ListBoards() ([]*models.Board, error) {
	var b []*models.Board
	op := &gitlab.ProjectListOptions{}
	op.Page = "1"
	op.PerPage = "100"
	op.Archived = "false"

	r, err := ds.client.ListProjects(op)

	if err != nil {
		return nil, err
	}

	for _, item := range r {
		b = append(b, mapBoardFromGitlab(item))
	}

	return b, nil
}

// ItemBoard returns board item
func (ds GitLabDataSource) ItemBoard(board_id string) (*models.Board, error) {
	var b *models.Board
	r, err := ds.client.ItemProject(board_id)

	if err, ok := err.(gitlab.ResponseError); ok {
		return nil, models.ReceivedDataErr{err.Error(), err.StatusCode}
	}

	if err != nil {
		return nil, err
	}
	b = mapBoardFromGitlab(r)

	return b, nil
}

// ConfigureBoard creates default stages for board
func (ds GitLabDataSource) ConfigureBoard(f *models.BoardRequest) (int, error) {
	b, err := ds.ItemBoard(f.BoardId)

	if err != nil {
		return 0, err
	}

	for _, stage := range defaultStages {
		_, res, err := ds.client.CreateLabel(fmt.Sprintf("%d", b.Id), &gitlab.LabelRequest{
			Name:  stage,
			Color: "#F5F5F5",
		})

		if err != nil {
			return res.StatusCode, err
		}
	}

	return 0, nil
}

// CreateConnectBoard connects other board to current
// for show all cards from other boards
func (ds GitLabDataSource) CreateConnectBoard(BoardID, ConnectBoardID string) (int, error) {
	current, err := ds.ItemBoard(BoardID)

	if err != nil {
		return 0, err
	}
	con, err := ds.ItemBoard(ConnectBoardID)

	if err != nil {
		return 0, err
	}

	_, err = ds.db.LPush(fmt.Sprintf("boards:%d:connect", current.Id), fmt.Sprintf("%d", con.Id)).Result()

	if err != nil {
		return 0, err
	}

	return 0, nil
}

// ListConnectBoard return list connect board for current board
func (ds GitLabDataSource) ListConnectBoard(boardID string) ([]*models.Board, int, error) {
	b := []*models.Board{}

	boards, err := ds.db.LRange(fmt.Sprintf("boards:%s:connect", boardID), 0, 100).Result()

	if err != nil {
		return nil, 0, err
	}

	for _, board := range boards {
		item, _ := ds.ItemBoard(board)
		b = append(b, item)
	}
	return b, 0, nil
}

// DeleteConnectBoard deletes from connected board list board
func (ds GitLabDataSource) DeleteConnectBoard(boardID, ConnectBoardID string) (int, error) {

	_, err := ds.db.LRem(fmt.Sprintf("boards:%s:connect", boardID), 0, ConnectBoardID).Result()

	if err != nil {
		return 0, err
	}

	return 0, nil
}

// mapBoardFromGitlab transforms board from gitlab project to kanban
func mapBoardFromGitlab(r *gitlab.Project) *models.Board {
	return &models.Board{
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
func mapNamespaceFromGitlab(n *gitlab.Namespace) *models.Namespace {
	if n == nil {
		return nil
	}
	return &models.Namespace{
		Id:     n.Id,
		Name:   n.Name,
		Avatar: mapAvatarFromGitlab(n.Avatar),
	}
}

// mapAvatarFromGitlab transform gitlab avatar to kanban avatar
func mapAvatarFromGitlab(n *gitlab.Avatar) *models.Avatar {
	if n == nil {
		return nil
	}
	return &models.Avatar{
		Url: n.Url,
	}
}
