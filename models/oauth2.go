package models

import (
	"gitlab.com/leanlabsio/kanban/modules/gitlab"
	"golang.org/x/oauth2"
)

//AuthCodeUrl return url for redirect user for oauth.
func AuthCodeURL(p string) string {
	switch p {
	case "gitlab":
		return gitlab.AuthCodeURL()
	}

	return ""
}

// Exchange converts an authorization code into a token.
func Exchange(p string, c string) (*oauth2.Token, error) {
	switch p {
	case "gitlab":
		return gitlab.Exchange(c)
	}

	return &oauth2.Token{}, nil
}

// UserOauthSignIn is
func UserOauthSignIn(provider string, tok *oauth2.Token) (*User, error) {

	m := map[string]*Credential{
		provider: &Credential{
			Token:        tok,
			PrivateToken: "",
		},
	}

	user := &User{
		Credential: m,
	}

	return LoadByToken(user, provider)
}
