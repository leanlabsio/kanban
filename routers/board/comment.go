package board

import (
	"gitlab.com/leanlabsio/kanban/models"
	"gitlab.com/leanlabsio/kanban/modules/middleware"
	"net/http"
)

// ListComments gets a list of comment on board and card
// accessible by the authenticated user.
func ListComments(ctx *middleware.Context) {
	boards, err := models.ListComments(ctx.User, ctx.Provider, ctx.Query("project_id"), ctx.Query("issue_id"))

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

// CreateComment creates new kanban comment
func CreateComment(ctx *middleware.Context, form models.CommentRequest) {
	com, code, err := models.CreateComment(ctx.User, ctx.Provider, &form)

	if err != nil {
		ctx.JSON(code, &models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &models.Response{
		Data: com,
	})
}
