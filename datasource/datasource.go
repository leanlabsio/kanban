package datasource

import (
	"gitlab.com/leanlabsio/kanban/models"
)

type DataSource interface {
	CardSource
	BoardSource
	CommentSource
	UserSource
	LabelSource
	MilestoneSource
}

type CardSource interface {
	ListCards(project_id string) ([]*models.Card, error)
	CreateCard(*models.CardRequest) (*models.Card, int, error)
	UpdateCard(*models.CardRequest) (*models.Card, int, error)
	DeleteCard(*models.CardRequest) (*models.Card, int, error)
}

type BoardSource interface {
	ListBoards() ([]*models.Board, error)
	ItemBoard(board_id string) (*models.Board, error)
	ConfigureBoard(*models.BoardRequest) (int, error)
}

type CommentSource interface {
	ListComments(project_id, card_id string) ([]*models.Comment, error)
	CreateComment(*models.CommentRequest) (*models.Comment, int, error)
}

type UserSource interface {
	ListMembers(board_id string) ([]*models.User, error)
}

type LabelSource interface {
	ListLabels(board_id string) ([]*models.Label, error)
}

type MilestoneSource interface {
	ListMilestones(board_id string) ([]*models.Milestone, error)
	CreateMilestone(*models.MilestoneRequest) (*models.Milestone, int, error)
}
