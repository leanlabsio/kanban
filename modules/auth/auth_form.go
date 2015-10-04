package auth

type Oauth2 struct {
	Code     string `json:"code"`
	Provider string `json:"provider"`
}
