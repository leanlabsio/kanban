package board

import (
	"fmt"
	"net/http"

	"gitlab.com/leanlabsio/kanban/models"
	"gitlab.com/leanlabsio/kanban/modules/middleware"
)

// ListCards gets a list of card on board accessible by the authenticated user.
func ListCards(ctx *middleware.Context) {
	pr := ctx.Query("project_id")

	proj, _, err := ctx.DataSource.ListConnectBoard(pr)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	current, err := ctx.DataSource.ItemBoard(pr)

	if err != nil {
		ctx.JSON(http.StatusNotFound, &models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	proj = append(proj, current)

	var cards, c []*models.Card

	for _, p := range proj {
		c, err = ctx.DataSource.ListCards(p)
		cards = append(cards, c...)
	}

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
	card, code, err := ctx.DataSource.CreateCard(&form)

	if err != nil {
		ctx.JSON(code, &models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	card.BoardID = ctx.Params(":board")

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
	card, code, err := ctx.DataSource.UpdateCard(&form)

	if err != nil {
		ctx.JSON(code, &models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	card.BoardID = ctx.Params(":board")

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
	card, code, err := ctx.DataSource.DeleteCard(&form)

	if err != nil {
		ctx.JSON(code, &models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	card.BoardID = ctx.Params(":board")

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
	card, code, err := ctx.DataSource.UpdateCard(&form)

	if err != nil {
		ctx.JSON(code, &models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	card.BoardID = ctx.Params(":board")

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
			ctx.DataSource.CreateComment(&com)
		}()
	}
}

// ChangeProjectForCard locate card to another project
func ChangeProjectForCard(ctx *middleware.Context, form models.CardRequest) {
	oldCrard := models.Card{
		Id:        form.CardId,
		ProjectId: form.ProjectId,
		BoardID:   ctx.Params(":board"),
	}

	card, code, err := ctx.DataSource.ChangeProjectForCard(&form, ctx.Params(":projectId"))

	if err != nil {
		ctx.JSON(code, &models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.Broadcast(oldCrard.RoutingKey(), &models.Response{
		Data:  oldCrard,
		Event: "card.delete",
	})

	card.BoardID = ctx.Params(":board")

	ctx.Broadcast(card.RoutingKey(), &models.Response{
		Data:  card,
		Event: "card.create",
	})

	ctx.JSON(http.StatusOK, &models.Response{
		Data: card,
	})
}
