package models

import (
	"errors"
)

// UserSignIn is
func UserSignIn(uname, pass string) (*User, error) {
	u, err := LoadUserByUsername(uname)

	if err != nil {
		return nil, err
	}

	if !u.ValidatePassword(pass) {
		return nil, errors.New("Invalid username or password")
	}

	return u, nil
}

// UserSignUp is
func UserSignUp(uname, email, pass, token, provider string) (*User, error) {
	cr := map[string]*Credential{
		provider: &Credential{
			Token: nil,
			PrivateToken: token,
		},
	}

	u := &User{
		Name:       uname,
		Username:   uname,
		Email:      email,
		Passwd:     pass,
		Credential: cr,
	}

	return CreateUser(u)
}
