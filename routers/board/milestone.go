package board

import (
	"gitlab.com/leanlabsio/kanban/models"
	"gitlab.com/leanlabsio/kanban/modules/middleware"
	"net/http"
)

// ListMilestones gets a list of milestone on board accessible by the authenticated user.
func ListMilestones(ctx *middleware.Context) {
	labels, err := models.ListMilestones(ctx.User, ctx.Provider, ctx.Query("project_id"))

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, &models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &models.Response{
		Data: labels,
	})
}
