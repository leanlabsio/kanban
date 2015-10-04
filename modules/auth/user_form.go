package auth

type SignIn struct {
	Uname string `json:"_username"`
	Pass  string `json:"_password"`
}

type SignUp struct {
	Email string `json:"_email"`
	Pass  string `json:"_password"`
	Token string `json:"_token"`
	Uname string `json:"_username"`
}

type ResponseAuth struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}
