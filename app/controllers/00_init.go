package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const (
	// FlashError 错误提示
	FlashError = "FlashError"
)

var (
	// CookieTag cookie key
	CookieTag      string
	cookieSecure   bool
	cookieHTTPOnly bool
	cookieExprHour int
	cookieDomain   string
)

// Init 初始化
func Init() {
	CookieTag = viper.GetString("cookie.tag")
	if CookieTag == "" {
		CookieTag = "_t"
	}
	cookieDomain = viper.GetString("cookie.domain")
	cookieSecure = viper.GetBool("cookie.secure")
	cookieHTTPOnly = viper.GetBool("cookie.http_only")
	cookieExprHour = viper.GetInt("cookie.expr_hour") * 3600
}

// setJustCookie 设置 cookie
// 默认的基础设置
func setJustCookie(c *gin.Context, token string) {
	c.SetCookie(CookieTag, token, cookieExprHour, "", cookieDomain, cookieSecure, cookieHTTPOnly)
}

// 失效 cookie
func expireCookie(c *gin.Context) {
	c.SetCookie(CookieTag, "", 0, "", cookieDomain, cookieSecure, cookieHTTPOnly)

}
