package board

import (
	"gitlab.com/leanlabsio/kanban/models"
	"gitlab.com/leanlabsio/kanban/modules/middleware"
	"net/http"
)

// ListBoards gets a list of board accessible by the authenticated user.
func ListBoards(ctx *middleware.Context) {
	boards, err := models.ListBoards(ctx.User, ctx.Provider)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, &models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &models.Response{
		Data: boards,
	})
}

// ItemBoard gets a specific board, identified by project ID or
// NAMESPACE/BOARD_NAME, which is owned by the authenticated user.
func ItemBoard(ctx *middleware.Context) {
	board, err := models.ItemBoard(ctx.User, ctx.Provider, ctx.Query("project_id"))

	if err != nil {
		if err, ok := err.(models.ReceivedDataErr); ok {
			ctx.JSON(err.StatusCode, &models.ResponseError{
				Success: false,
				Message: err.Error(),
			})
		}
		ctx.JSON(http.StatusInternalServerError, &models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &models.Response{
		Data: board,
	})
}

func Configure(ctx *middleware.Context, form models.BoardRequest) {
	status, err := models.ConfigureBoard(ctx.User, ctx.Provider, &form)

	if err != nil {
		ctx.JSON(status, &models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &models.Response{})
}
