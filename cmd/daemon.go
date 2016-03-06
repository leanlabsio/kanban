package cmd

import (
	"github.com/go-macaron/bindata"
	"github.com/go-macaron/binding"
	"github.com/leanlabsio/sockets"
	"github.com/spf13/cobra"
	"gitlab.com/leanlabsio/kanban/templates"
	"gitlab.com/leanlabsio/kanban/web"
	"gitlab.com/leanlabsio/kanban/ws"
	"gopkg.in/macaron.v1"
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
	DaemonCmd.Flags().Bool(
		"enable-signup",
		true,
		"Enable signup",
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
				AssetInfo:  templates.AssetInfo,
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
				AssetInfo:  web.AssetInfo,
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
				AssetInfo:  web.AssetInfo,
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
				AssetInfo:  web.AssetInfo,
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
		m.Get("/labels/:project", middleware.Datasource(), board.ListLabels)
		m.Put("/labels/:project", middleware.Datasource(), binding.Json(models.LabelRequest{}), board.EditLabel)
		m.Delete("/labels/:project/:label", middleware.Datasource(), board.DeleteLabel)
		m.Post("/labels/:project", middleware.Datasource(), binding.Json(models.LabelRequest{}), board.CreateLabel)

		m.Get("/boards", middleware.Datasource(), board.ListBoards)
		m.Post("/boards/configure", middleware.Datasource(), binding.Json(models.BoardRequest{}), board.Configure)

		m.Get("/board", middleware.Datasource(), board.ItemBoard)

		m.Get("/cards", middleware.Datasource(), board.ListCards)
		m.Combo("/milestones").
			Get(middleware.Datasource(), board.ListMilestones).
			Post(middleware.Datasource(), binding.Json(models.MilestoneRequest{}), board.CreateMilestone)

		m.Get("/users", middleware.Datasource(), board.ListMembers)
		m.Combo("/comments").
			Get(middleware.Datasource(), board.ListComments).
			Post(middleware.Datasource(), binding.Json(models.CommentRequest{}), board.CreateComment)

		m.Combo("/card").
			Post(middleware.Datasource(), binding.Json(models.CardRequest{}), board.CreateCard).
			Put(middleware.Datasource(), binding.Json(models.CardRequest{}), board.UpdateCard).
			Delete(middleware.Datasource(), binding.Json(models.CardRequest{}), board.DeleteCard)

		m.Put("/card/move", middleware.Datasource(), binding.Json(models.CardRequest{}), board.MoveToCard)

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
