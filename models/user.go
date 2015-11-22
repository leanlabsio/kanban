package models

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"gitlab.com/leanlabsio/kanban/modules/gitlab"
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
	Passwd     string
	Salt       string
	Email      string
}

type UserNotFound struct {
	msg string
}

type Credential struct {
	PrivateToken string `json:"private_access_token"`
	Token        *oauth2.Token
}

type ResponseUser struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Token   string `json:"token"`
	Success bool   `json:"success"`
}

// LoadUserByUsername is
func LoadUserByUsername(uname string) (*User, error) {
	cmd := c.HGetAllMap(fmt.Sprintf("kanban:users:%s", strings.ToLower(uname)))
	val, err := cmd.Result()

	if err != nil {
		panic(fmt.Sprintf("Redis unrecoverable error: %s", err))
	}

	if len(val) == 0 {
		return nil, &UserNotFound{msg: "User not found"}
	}

	u := User{
		Username: val["username"],
		Passwd:   val["password"],
		Email:    val["email"],
		Name:     val["name"],
	}

	if val["credentials"] != "" {
		var cr map[string]*Credential
		err = json.Unmarshal([]byte(val["credentials"]), &cr)

		if err != nil {
			return nil, err
		}
		u.Credential = cr
	}

	return &u, nil
}

// LoadByToken is
func LoadByToken(u *User, provider string) (*User, error) {
	var user *User
	switch provider {
	case "gitlab":
		c := gitlab.NewContext(u.Credential["gitlab"].Token, u.Credential["gitlab"].PrivateToken)
		r, err := c.CurrentUser()

		if err != nil {
			return nil, err
		}
		user = mapUserFromGitlab(r)
		user.Credential = u.Credential
		user.Credential[provider].PrivateToken = r.PrivateToken
	}

	return user, nil
}

// CreateUser creates new user
func CreateUser(u *User) (*User, error) {
	usr, _ := LoadUserByUsername(u.Username)

	if usr != nil {
		return nil, errors.New("User already exists")
	}

	u.EncodePasswd()
	return saveUser(u)
}

// UpdateUser updates user's information.
func UpdateUser(u *User) (*User, error) {
	user, err := LoadUserByUsername(u.Username)
	if err != nil {
		return user, err
	}

	if user == nil {
		return user, errors.New(fmt.Sprintf("User with username %s does not exists", u.Username))
	}

	user.Credential = u.Credential

	return saveUser(user)
}

// saveUser saved user's information.
func saveUser(u *User) (*User, error) {
	val, err := json.Marshal(u.Credential)
	if err != nil {
		return nil, err
	}

	_, err = c.HMSet("kanban:users:"+strings.ToLower(u.Username),
		"credentials", string(val),
		"name", u.Name,
		"username", u.Username,
		"password", u.Passwd,
		"email", u.Email,
	).Result()

	if err != nil {
		return nil, err
	}

	return &User{
		Name:       u.Name,
		Username:   u.Username,
		Credential: u.Credential,
		Passwd:     u.Passwd,
		Email:      u.Email,
	}, nil
}

// ListMembers is
func ListMembers(u *User, provider, board_id string) ([]*User, error) {
	var mem []*User
	switch provider {
	case "gitlab":
		c := gitlab.NewContext(u.Credential["gitlab"].Token, u.Credential["gitlab"].PrivateToken)
		r, err := c.ListProjectMembers(board_id, &gitlab.ListOptions{
			Page:    "1",
			PerPage: "100",
		})

		if err != nil {
			return nil, err
		}

		b, err := ItemBoard(u, provider, board_id)
		fmt.Printf("%+v", b)

		if err != nil {
			return nil, err
		}

		exist := make(map[string]bool)

		if b.Owner == nil {
			u, _ := c.ListGroupMembers(fmt.Sprintf("%d", b.Namespace.Id), &gitlab.ListOptions{})

			for _, item := range u {
				exist[item.Username] = true
				mem = append(mem, mapUserFromGitlab(item))
			}
		}

		for _, item := range r {
			if !exist[item.Username] {
				mem = append(mem, mapUserFromGitlab(item))
			}
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
		return token.SignedString([]byte(viper.GetString("security.secret_key")))
	}

	return u.Token.SigningString()
}

// ValidatePassword checks if given password matches the one belongs to the user.
func (u *User) ValidatePassword(pass string) bool {
	newUser := &User{Passwd: pass, Salt: ""}
	newUser.EncodePasswd()
	return u.Passwd == newUser.Passwd
}

// EncodePasswd encodes password to safe format.
func (u *User) EncodePasswd() {
	hash := []byte(u.Passwd)
	dig := sha512.Sum512(hash)
	for i := 1; i < 5000; i++ {
		dig = sha512.Sum512(append(dig[:], hash[:]...))
	}
	newPasswd := base64.StdEncoding.EncodeToString(dig[:])
	u.Passwd = fmt.Sprintf("%s", newPasswd)
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

func (e *UserNotFound) Error() string {
	return e.msg
}
