package model

type Oauth2 struct {
	Code     string `json:"code"`
	Provider string `json:"provider"`
}

type KbAuth struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}
