package board

import (
	"gitlab.com/kanban/kanban/models"
	"gitlab.com/kanban/kanban/modules/middleware"
	"net/http"
)

func ListCards(ctx *middleware.Context) {
	cards, err := models.ListCards(ctx.User, ctx.Provider, ctx.Query("project_id"))

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, &models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &models.Response{
		Data: cards,
	})
}

func CreateCard(ctx *middleware.Context, form models.CardRequest) {
	card, code, err := models.CreateCard(ctx.User, ctx.Provider, &form)

	if err != nil {
		ctx.JSON(code, &models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &models.Response{
		Data: card,
	})
}

func UpdateCard(ctx *middleware.Context, form models.CardRequest) {
	card, code, err := models.UpdateCard(ctx.User, ctx.Provider, &form)

	if err != nil {
		ctx.JSON(code, &models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &models.Response{
		Data: card,
	})
}

func DeleteCard(ctx *middleware.Context, form models.CardRequest) {
	card, code, err := models.DeleteCard(ctx.User, ctx.Provider, &form)

	if err != nil {
		ctx.JSON(code, &models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &models.Response{
		Data: card,
	})
}

func MoveToCard(ctx *middleware.Context, form models.CardRequest) {
	card, code, err := models.UpdateCard(ctx.User, ctx.Provider, &form)

	// Todo implement method for add comments
	if err != nil {
		ctx.JSON(code, &models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &models.Response{
		Data: card,
	})
}
