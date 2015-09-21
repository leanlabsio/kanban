package board

import (

	"gitlab.com/kanban/kanban/modules/middleware"
	"net/http"
	"gitlab.com/kanban/kanban/models"
)

func ListBoards(ctx *middleware.Context) {
	boards, err := models.ListBoards(ctx.User, ctx.Provider)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, &models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, &models.Response{
		Data: boards,
	})
}

// ItemBoard is
func ItemBoard(ctx *middleware.Context) {
	board, err := models.ItemBoard(ctx.User, ctx.Provider, ctx.Query("project_id"))

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, &models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, &models.Response{
		Data: board,
	})
}