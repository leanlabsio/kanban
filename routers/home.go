package routers

import (
	"github.com/Unknwon/macaron"
	"github.com/spf13/viper"
)

// Home returns main page
func Home(ctx *macaron.Context) {
	ctx.Data["Version"] = viper.GetString("version")
	ctx.Data["GitlabHost"] = viper.GetString("gitlab.url")
	ctx.HTML(200, "templates/index")
}
