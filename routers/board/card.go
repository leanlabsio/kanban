package board
import (
	"gitlab.com/kanban/kanban/modules/middleware"
	"gitlab.com/kanban/kanban/models"
	"net/http"
)

func ListCards(ctx *middleware.Context) {
	boards, err := models.ListCards(ctx.User, ctx.Provider, ctx.Query("project_id"))

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

