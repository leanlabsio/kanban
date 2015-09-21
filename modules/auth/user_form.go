package auth

type SignIn struct {
	Login string `json:"login"`
	Pass  string `json:"password"`
}

type SignUp struct {
	Login string `json:"login"`
	Email string `json:"email"`
	Pass  string `json:"password"`
	Token string `json:"token"`
}

type ResponseAuth struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}
