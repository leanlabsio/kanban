package gitlab

import (
	"github.com/Unknwon/macaron"
	"net/http"
)

type ApiErr struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Message string `json: "message"`
}

func (e *ApiErr) Error() string {
	return e.Message
}

// Generate error for client
func (g *ApiGitlab) SendError(ctx *macaron.Context, err error) {
	ctx.JSON(http.StatusBadRequest, ApiErr{
		Success: false,
		Message: err.Error(),
	})
}
