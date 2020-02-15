package mailer

import (
	"github.com/spf13/viper"
)

var (
	emailAddress   string
	emailAliasName string
	emailPassword  string
	emailHost      string
	emailHostPort  int
)

// Init 初始化
func Init() {
	emailAddress = viper.GetString("mailer.address")
	emailAliasName = viper.GetString("mailer.alias_name")
	emailPassword = viper.GetString("mailer.password")
	emailHost = viper.GetString("mailer.host")
	emailHostPort = viper.GetInt("mailer.port")
}
