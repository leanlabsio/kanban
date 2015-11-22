package user

import (
	"fmt"
	"github.com/Unknwon/macaron"
	"gitlab.com/leanlabsio/kanban/models"
	"gitlab.com/leanlabsio/kanban/modules/auth"
	"gitlab.com/leanlabsio/kanban/modules/middleware"
	"log"
	"net/http"
)

// OauthUrl redirects to url for authorisation
func OauthUrl(ctx *middleware.Context) {
	ctx.Redirect(models.AuthCodeURL(ctx.Query("provider")))
}

// OauthLogin logins with gitlab and get access token
func OauthLogin(ctx *middleware.Context, form auth.Oauth2) {
	tok, err := models.Exchange(form.Provider, form.Code)

	if err != nil {
		log.Printf("%s", err.Error())
		ctx.JSON(http.StatusBadRequest, models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	user, err := models.UserOauthSignIn(form.Provider, tok)

	if err != nil {
		log.Printf("%s", err.Error())
		ctx.JSON(http.StatusBadRequest, models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	user, err = models.LoadByToken(user, ctx.Provider)
	if err != nil {
		log.Printf("%s", err.Error())
		ctx.JSON(http.StatusBadRequest, models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	user.Username = fmt.Sprintf("%s_%s", user.Username, ctx.Provider)
	_, err = models.UpdateUser(user)

	// todo add validation by oauth provider
	if err != nil {
		user, err = models.CreateUser(user)
	}

	if err != nil {
		log.Printf("%s", err.Error())
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	tokens, err := user.SignedString()

	if err != nil {
		log.Printf("%s", err.Error())
		ctx.JSON(http.StatusBadRequest, models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, auth.ResponseAuth{
		Success: true,
		Token:   tokens,
	})
}

// SignIn registers with user data
func SignIn(ctx *macaron.Context, form auth.SignIn) {
	u, err := models.UserSignIn(form.Uname, form.Pass)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	tokens, _ := u.SignedString()

	ctx.JSON(http.StatusOK, auth.ResponseAuth{
		Success: true,
		Token:   tokens,
	})
}

// SignUp sing ups with data
func SignUp(ctx *middleware.Context, form auth.SignUp) {
	u, err := models.UserSignUp(form.Uname, form.Email, form.Pass, form.Token, ctx.Provider)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	tokens, _ := u.SignedString()

	ctx.JSON(http.StatusOK, auth.ResponseAuth{
		Success: true,
		Token:   tokens,
	})
}

// OauthHandler handles request from other services
func OauthHandler(ctx *middleware.Context) {
	ctx.HTML(200, "templates/oauth")
}
