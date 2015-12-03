package routers

import (
	"github.com/spf13/viper"
	"gopkg.in/macaron.v1"
)

// Home returns main page
func Home(ctx *macaron.Context) {
	ctx.Data["Version"] = viper.GetString("version")
	ctx.Data["GitlabHost"] = viper.GetString("gitlab.url")
	ctx.HTML(200, "templates/index")
}
