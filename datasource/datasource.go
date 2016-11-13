package datasource

import "gitlab.com/leanlabsio/kanban/models"

type DataSource interface {
	CardSource
	BoardSource
	CommentSource
	UserSource
	LabelSource
	MilestoneSource
	FileService
}

type CardSource interface {
	ListCards(*models.Board) ([]*models.Card, error)
	CreateCard(*models.CardRequest) (*models.Card, int, error)
	UpdateCard(*models.CardRequest) (*models.Card, int, error)
	DeleteCard(*models.CardRequest) (*models.Card, int, error)
	ChangeProjectForCard(form *models.CardRequest, ToProjectID string) (*models.Card, int, error)
}

type BoardSource interface {
	ListBoards() ([]*models.Board, error)
	ListStarredBoards() ([]*models.Board, error)
	ItemBoard(boardID string) (*models.Board, error)
	ConfigureBoard(*models.BoardRequest) (int, error)
	CreateConnectBoard(BoardID, ConnectedBoardID string) (int, error)
	ListConnectBoard(BoardID string) ([]*models.Board, int, error)
	DeleteConnectBoard(BoardID, ConnectBoardID string) (int, error)
}

type CommentSource interface {
	ListComments(projectID string, cardID string) ([]*models.Comment, error)
	CreateComment(*models.CommentRequest) (*models.Comment, int, error)
}

type UserSource interface {
	ListMembers(boardID string) ([]*models.User, error)
}

type LabelSource interface {
	ListLabels(boardID string) ([]*models.Label, error)
	EditLabel(projectID string, req *models.LabelRequest) (*models.Label, error)
	DeleteLabel(projectID, name string) (*models.Label, error)
	CreateLabel(projectID string, req *models.LabelRequest) (*models.Label, error)
}

type MilestoneSource interface {
	ListMilestones(boardID string) ([]*models.Milestone, error)
	CreateMilestone(*models.MilestoneRequest) (*models.Milestone, int, error)
}

// FileService represents uploaded file
type FileService interface {
	UploadFile(boardID string, file models.UploadForm) (*models.File, error)
}
