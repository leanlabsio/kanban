package cmd

import (
	"github.com/Unknwon/macaron"
	"github.com/macaron-contrib/bindata"
	"github.com/macaron-contrib/binding"
	"github.com/macaron-contrib/sockets"
	"github.com/spf13/cobra"
	"gitlab.com/leanlabsio/kanban/templates"
	"gitlab.com/leanlabsio/kanban/web"
	"gitlab.com/leanlabsio/kanban/ws"
	"log"
	"net/http"

	"gitlab.com/leanlabsio/kanban/modules/auth"
	"gitlab.com/leanlabsio/kanban/modules/setting"

	"gitlab.com/leanlabsio/kanban/models"
	"gitlab.com/leanlabsio/kanban/modules/middleware"
	"gitlab.com/leanlabsio/kanban/routers"
	"gitlab.com/leanlabsio/kanban/routers/board"
	"gitlab.com/leanlabsio/kanban/routers/user"

	"github.com/spf13/viper"
)

// DaemonCmd is implementation of command to run application in daemon mode
var DaemonCmd = cobra.Command{
	Use:   "server",
	Short: "Starts LeanLabs Kanban board application",
	Long: `Start LeanLabs Kanban board application.

Please refer to http://kanban.leanlabs.io/documentation/Home for full documentation.

Report bugs to <support@leanlabs.io> or https://gitter.im/leanlabsio/kanban.
        `,
	Run: daemon,
}

func init() {
	DaemonCmd.Flags().String(
		"server-listen",
		"0.0.0.0:80",
		"IP:PORT to listen on",
	)
	DaemonCmd.Flags().String(
		"server-hostname",
		"http://localhost",
		"URL on which Leanlabs Kanban will be reachable",
	)
	DaemonCmd.Flags().String(
		"security-secret",
		"qwerty",
		"This string is used to generate user auth tokens",
	)
	DaemonCmd.Flags().String(
		"gitlab-url",
		"https://gitlab.com",
		"Your GitLab host URL",
	)
	DaemonCmd.Flags().String(
		"gitlab-client",
		"qwerty",
		"Your GitLab OAuth client ID",
	)
	DaemonCmd.Flags().String(
		"gitlab-secret",
		"qwerty",
		"Your GitLab OAuth client secret key",
	)
	DaemonCmd.Flags().String(
		"redis-addr",
		"127.0.0.1:6379",
		"Redis server address - IP:PORT",
	)
	DaemonCmd.Flags().String(
		"redis-password",
		"",
		"Redis server password, empty string if none",
	)
	DaemonCmd.Flags().Int64(
		"redis-db",
		0,
		"Redis server database numeric index, from 0 to 16",
	)
}

func daemon(c *cobra.Command, a []string) {
	m := macaron.New()
	setting.NewContext(c)
	err := models.NewEngine()

	m.Use(middleware.Contexter())
	m.Use(macaron.Recovery())
	m.Use(macaron.Logger())
	m.Use(macaron.Renderer(
		macaron.RenderOptions{
			Directory: "templates",
			TemplateFileSystem: bindata.Templates(bindata.Options{
				Asset:      templates.Asset,
				AssetDir:   templates.AssetDir,
				AssetNames: templates.AssetNames,
				Prefix:     "",
			}),
		},
	))
	m.Use(macaron.Static("web/images",
		macaron.StaticOptions{
			Prefix: "images",
			FileSystem: bindata.Static(bindata.Options{
				Asset:      web.Asset,
				AssetDir:   web.AssetDir,
				AssetNames: web.AssetNames,
				Prefix:     "web/images",
			}),
		},
	))
	m.Use(macaron.Static("web/template",
		macaron.StaticOptions{
			Prefix: "template",
			FileSystem: bindata.Static(bindata.Options{
				Asset:      web.Asset,
				AssetDir:   web.AssetDir,
				AssetNames: web.AssetNames,
				Prefix:     "web/template",
			}),
		},
	))

	m.Use(macaron.Static("web",
		macaron.StaticOptions{
			FileSystem: bindata.Static(bindata.Options{
				Asset:      web.Asset,
				AssetDir:   web.AssetDir,
				AssetNames: web.AssetNames,
				Prefix:     "web",
			}),
			Prefix: viper.GetString("version"),
		},
	))

	m.Get("/assets/html/user/views/oauth.html", user.OauthHandler)
	m.Combo("/api/oauth").
		Get(user.OauthUrl).
		Post(binding.Json(auth.Oauth2{}), user.OauthLogin)

	m.Post("/api/login", binding.Json(auth.SignIn{}), user.SignIn)
	m.Post("/api/register", binding.Json(auth.SignUp{}), user.SignUp)
	m.Group("/api", func() {
		m.Get("/boards", board.ListBoards)
		m.Post("/boards/configure", binding.Json(models.BoardRequest{}), board.Configure)

		m.Get("/board", board.ItemBoard)
		m.Get("/labels", board.ListLabels)
		m.Get("/cards", board.ListCards)
		m.Get("/milestones", board.ListMilestones)
		m.Get("/users", board.ListMembers)
		m.Combo("/comments").
			Get(board.ListComments).
			Post(binding.Json(models.CommentRequest{}), board.CreateComment)

		m.Combo("/card").
			Post(binding.Json(models.CardRequest{}), board.CreateCard).
			Put(binding.Json(models.CardRequest{}), board.UpdateCard).
			Delete(binding.Json(models.CardRequest{}), board.DeleteCard)

		m.Put("/card/move", binding.Json(models.CardRequest{}), board.MoveToCard)

	}, middleware.Auther())
	m.Get("/*", routers.Home)
	m.Get("/ws/", sockets.Messages(), ws.ListenAndServe)

	listen := viper.GetString("server.listen")
	log.Printf("Listen: %s", listen)
	err = http.ListenAndServe(listen, m)

	if err != nil {
		log.Fatalf("Failed to start: %s", err)
	}
}
