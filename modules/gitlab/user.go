package gitlab

type User struct {
	Id        int64  `json:"id"`
	Name      string `json:"name,omitempty"`
	AvatarUrl string `json:"avatar_url,nil,omitempty"`
	State     string `json:"state,omitempty"`
	Username  string `json:"username,omitempty"`
	WebUrl    string `json:"web_url,omitempty"`
}
