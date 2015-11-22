package setting

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

// NewContext created new context for settings
func NewContext(c *cobra.Command) {
	repl := strings.NewReplacer(".", "_")

	viper.SetEnvPrefix("kanban")
	viper.SetEnvKeyReplacer(repl)

	viper.SetDefault("server.listen", "0.0.0.0:80")
	viper.BindEnv("server.listen")
	if c.Flags().Lookup("server-listen").Changed {
		viper.BindPFlag("server.listen", c.Flags().Lookup("server-listen"))
	}

	viper.SetDefault("server.hostname", "http://localhost")
	viper.BindEnv("server.hostname")
	if c.Flags().Lookup("server-hostname").Changed {
		viper.BindPFlag("server.hostname", c.Flags().Lookup("server-hostname"))
	}

	viper.SetDefault("security.secret", "qwerty")
	viper.BindEnv("security.secret")
	if c.Flags().Lookup("security-secret").Changed {
		viper.BindPFlag("security.secret", c.Flags().Lookup("security-secret"))
	}

	viper.SetDefault("gitlab.url", "https://gitlab.com")
	viper.BindEnv("gitlab.url")
	if c.Flags().Lookup("gitlab-url").Changed {
		viper.BindPFlag("gitlab.url", c.Flags().Lookup("gitlab-url"))
	}

	viper.SetDefault("gitlab.client", "qwerty")
	viper.BindEnv("gitlab.client")
	if c.Flags().Lookup("gitlab-client").Changed {
		viper.BindPFlag("gitlab.client", c.Flags().Lookup("gitlab-client"))
	}

	viper.SetDefault("gitlab.secret", "qwerty")
	viper.BindEnv("gitlab.secret")
	if c.Flags().Lookup("gitlab-secret").Changed {
		viper.BindPFlag("gitlab.secret", c.Flags().Lookup("gitlab-secret"))
	}

	viper.SetDefault("redis.addr", "127.0.0.1:6379")
	viper.BindEnv("redis.addr")
	if c.Flags().Lookup("redis-addr").Changed {
		viper.BindPFlag("redis.addr", c.Flags().Lookup("redis-addr"))
	}

	viper.SetDefault("redis.password", "")
	viper.BindEnv("redis.password")
	if c.Flags().Lookup("redis-password").Changed {
		viper.BindPFlag("redis.password", c.Flags().Lookup("redis-password"))
	}

	viper.SetDefault("redis.db", 0)
	viper.BindEnv("redis.db")
	if c.Flags().Lookup("redis-db").Changed {
		viper.BindPFlag("redis.db", c.Flags().Lookup("redis-db"))
	}

	viper.SetDefault("version", "1.4.2")
}
