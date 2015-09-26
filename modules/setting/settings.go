package setting

import (
	"log"
	"github.com/spf13/viper"
	"github.com/codegangsta/cli"
)

// NewContext created new context for settings
func NewContext(c *cli.Context) {
	viper.SetConfigName("config")
	viper.AddConfigPath("conf")
	viper.AddConfigPath(c.String("config"))
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("Fatal error config file: %s \n", err)
	}

	viper.Set("Version", c.App.Version)
	if "" != c.String("cache-host") {
		viper.Set("cache.host", c.String("cache-host"))
	}
	if "" != c.String("gitlab-client-id") {
		viper.Set("gitlab.oauth_client_id", c.String("gitlab-client-id"))
	}
	if "" != c.String("gitlab-client-secret") {
		viper.Set("gitlab.oauth_client_secret", c.String("gitlab-client-secret"))
	}

}
