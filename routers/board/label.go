package board

import (
	"gitlab.com/leanlabsio/kanban/models"
	"gitlab.com/leanlabsio/kanban/modules/middleware"
	"log"
	"net/http"
)

// ListLabels gets a list of label on board accessible by the authenticated user.
func ListLabels(ctx *middleware.Context) {
	labels, err := ctx.DataSource.ListLabels(ctx.Params(":project"))

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

// Edit label updates existing project label
func EditLabel(ctx *middleware.Context, form models.LabelRequest) {
	log.Printf("GOT LABEL req %+v", form)
	label, err := ctx.DataSource.EditLabel(ctx.Params(":project"), &form)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, &models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, &models.Response{
		Data: label,
	})
}

func CreateLabel(ctx *middleware.Context, form models.LabelRequest) {
	label, err := ctx.DataSource.CreateLabel(ctx.Params(":project"), &form)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, &models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, &models.Response{Data: label})
}

// DeleteLabel removes existing project label
func DeleteLabel(ctx *middleware.Context) {
	label, err := ctx.DataSource.DeleteLabel(ctx.Params(":project"), ctx.Params(":label"))

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, &models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, &models.Response{
		Data: label,
	})
}
