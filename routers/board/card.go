package board

import (
	"fmt"
	"gitlab.com/leanlabsio/kanban/models"
	"gitlab.com/leanlabsio/kanban/modules/middleware"
	"net/http"
)

// ListCards gets a list of card on board accessible by the authenticated user.
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

// CreateCard creates a new board card.
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

	ctx.Broadcast(card.RoutingKey(), &models.Response{
		Data:  card,
		Event: "card.create",
	})
}

// UpdateCard updates an existing board card.
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

	ctx.Broadcast(card.RoutingKey(), &models.Response{
		Data:  card,
		Event: "card.update",
	})
}

// DeleteCard closed an existing board card.
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

	ctx.Broadcast(card.RoutingKey(), &models.Response{
		Data:  card,
		Event: "card.delete",
	})
}

// MoveToCard updates an existing board card.
func MoveToCard(ctx *middleware.Context, form models.CardRequest) {
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

	ctx.Broadcast(card.RoutingKey(), &models.Response{
		Data:  card,
		Event: "card.move",
	})

	source := models.ParseLabelToStage(form.Stage["source"])
	dest := models.ParseLabelToStage(form.Stage["dest"])

	if source.Name != dest.Name {
		com := models.CommentRequest{
			CardId:    form.CardId,
			ProjectId: form.ProjectId,
			Body:      fmt.Sprintf("moved issue from **%s** to **%s**", source.Name, dest.Name),
		}

		go func() {
			models.CreateComment(ctx.User, ctx.Provider, &com)
		}()
	}
}
