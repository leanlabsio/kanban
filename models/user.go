package models

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"gitlab.com/kanban/kanban/modules/gitlab"
	"gitlab.com/kanban/kanban/modules/setting"
	"golang.org/x/oauth2"
	"strings"
	"time"
)

type User struct {
	Id         int64
	Name       string
	IsAdmin    bool
	Token      *jwt.Token
	Credential map[string]*Credential
	AvatarUrl  string
	State      string
	Username   string
}

type Credential struct {
	Token        *oauth2.Token
	PrivateToken string
}

type ResponseUser struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Token   string `json:"token"`
	Success bool   `json:"success"`
}

var (
	fields = []string{
		"username",
		"email",
		"password",
		"credentials",
		"role",
		"token",
	}
)

// LoadUserByUsername is
func LoadUserByUsername(uname string) (*User, error) {
	cmd := c.HMGet("users:"+strings.ToLower(uname), "credentials", "name", "username")
	val, err := cmd.Result()
	if err != nil {
		return nil, err
	}

	var cr map[string]*Credential

	err = json.Unmarshal([]byte(val[0].(string)), &cr)

	if err != nil {
		return nil, err
	}

	return &User{
		Name:       val[1].(string),
		Username:   val[2].(string),
		Credential: cr,
	}, nil

	return nil, nil
}

// LoadByToken is
func LoadByToken(u *User, provider string) (*User, error) {
	var user *User
	switch provider {
	case "gitlab":
		c := gitlab.NewContext(u.Credential["gitlab"].Token)
		r, err := c.CurrentUser()

		if err != nil {
			return nil, err
		}
		user = mapUserFromGitlab(r)
		user.Credential = u.Credential
	}

	return user, nil
}

// Create creates new user
func Create(u *User) (*User, error) {
	val, err := json.Marshal(u.Credential)
	if err != nil {
		return nil, err
	}
	_, err = c.HMSet("users:"+strings.ToLower(u.Username),
		"credentials", string(val),
		"name", u.Name,
		"username", u.Username,
	).Result()

	return &User{
		Name:       u.Name,
		Username:   u.Username,
		Credential: u.Credential,
	}, err
}

// ListMembers is
func ListMembers(u *User, provider, board_id string) ([]*User, error) {
	var mem []*User
	switch provider {
	case "gitlab":
		c := gitlab.NewContext(u.Credential["gitlab"].Token)
		r, err := c.ListProjectMembers(board_id)

		if err != nil {
			return nil, err
		}

		for _, item := range r {
			mem = append(mem, mapUserFromGitlab(item))
		}
	}

	return mem, nil
}

// SignedString returns user token for access
func (u *User) SignedString() (string, error) {
	if u.Token == nil {
		token := jwt.New(jwt.SigningMethodHS256)
		token.Claims["name"] = u.Username
		token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		return token.SignedString([]byte(setting.Cfg.Section("security").Key("SECRET_KEY").String()))
	}

	return u.Token.SigningString()
}

// mapUserFromGitlab mapped data from gitlab user to kanban user
func mapUserFromGitlab(u *gitlab.User) *User {
	if u == nil {
		return nil
	}
	return &User{
		Id:        u.Id,
		Name:      u.Name,
		Username:  u.Username,
		AvatarUrl: u.AvatarUrl,
		State:     u.State,
	}
}

// MarshalJSON returns the JSON encoding value.
func (u *User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(struct {
		Id        int64  `json:"id"`
		Name      string `json:"name,omitempty"`
		Username  string `json:"username,omitempty"`
		AvatarUrl string `json:"avatar_url,nil,omitempty"`
		State     string `json:"state,omitempty"`
	}{
		Id:        u.Id,
		Name:      u.Name,
		Username:  u.Username,
		AvatarUrl: u.AvatarUrl,
		State:     u.State,
	})
}
