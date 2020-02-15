package mailer

import (
	"fmt"

	"github.com/matcornic/hermes/v2"
	"github.com/spf13/viper"
)

var (
	emailAddress   string
	emailAliasName string
	emailPassword  string
	emailHost      string
	emailHostPort  int

	siteHost            string
	siteName            string
	hermesHeaderLogoURL string

	her hermes.Hermes
)

// Init 初始化
func Init() {
	emailAddress = viper.GetString("mailer.email.address")
	emailAliasName = viper.GetString("mailer.email.alias_name")
	emailPassword = viper.GetString("mailer.email.password")
	emailHost = viper.GetString("mailer.email.host")
	emailHostPort = viper.GetInt("mailer.email.port")

	siteHost = viper.GetString("core.site.host")
	siteName = viper.GetString("core.site.name")

	// 初始化 hermes，用来渲染邮件体
	// https://github.com/matcornic/hermes/blob/master/examples/main.go
	// Configure hermes by setting a theme and your product info
	her = hermes.Hermes{
		// Optional Theme
		Theme: new(hermes.Default),
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: siteName,
			Link: fmt.Sprintf("http://%s", siteHost),
			// Optional product logo
			Logo: viper.GetString("mailer.hermes.header_logo"),

			TroubleText: "如果不能点击按钮 '{ACTION}'，可以手动复制粘贴下面的链接到浏览器进行验证。",
		},
	}
}
