package user

import (
	"github.com/Unknwon/macaron"
	"gitlab.com/kanban/kanban/models"
	"gitlab.com/kanban/kanban/modules/auth"
	"gitlab.com/kanban/kanban/modules/middleware"
	"log"
	"net/http"
)

// OauthUrl redirects to url for authorisation
func OauthUrl(ctx *middleware.Context) {
	ctx.Redirect(models.AuthCodeURL(ctx.Query("provider")))
}

// Login with gitlab and get access token
func OauthLogin(ctx *middleware.Context, form auth.Oauth2) {
	tok, err := models.Exchange(form.Provider, form.Code)
	user, err := models.UserOauthSignIn(form.Provider, tok)

	if err != nil {
		log.Printf("%s", err.Error())
		ctx.JSON(http.StatusBadRequest, models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	user, err = models.Create(user)
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

// SignUp registers with user data
func SignUp(ctx *macaron.Context, form auth.SignIn) {
	User, err := models.UserSignIn(form.Login, form.Pass)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
	}

	tokens, _ := User.SignedString()

	ctx.JSON(http.StatusOK, auth.ResponseAuth{
		Success: true,
		Token:   tokens,
	})
}

// SignIn logins with data
func SignIn(ctx *middleware.Context, form auth.SignUp) {
	User, err := models.UserSignUp(form.Login, form.Email, form.Pass, form.Token)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
	}

	tokens, _ := User.SignedString()

	ctx.JSON(http.StatusOK, auth.ResponseAuth{
		Success: true,
		Token:   tokens,
	})
}

// OauthHandler handles request from other services
func OauthHandler(ctx *middleware.Context) {
	ctx.HTML(200, "templates/oauth")
}
