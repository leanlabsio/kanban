package board

import (
	"net/http"

	"gitlab.com/leanlabsio/kanban/models"
	"gitlab.com/leanlabsio/kanban/modules/middleware"
)

// UploadFile uploads file to datasource provider
func UploadFile(ctx *middleware.Context, f models.UploadForm) {
	res, err := ctx.DataSource.UploadFile(ctx.Params(":board"), f)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, &models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &models.Response{
		Data: res,
	})
}
