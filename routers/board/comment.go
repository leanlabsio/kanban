package board
import (
	"gitlab.com/kanban/kanban/modules/middleware"
	"gitlab.com/kanban/kanban/models"
	"net/http"
)

// ListComments returns list comments for cards
func ListComments(ctx *middleware.Context) {
	boards, err := models.ListComments(ctx.User, ctx.Provider, ctx.Query("project_id"), ctx.Query("issue_id"))

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