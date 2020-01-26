package router

import (
	"net/http"

	"github.com/chalvern/apollo/app/controllers"
	"github.com/chalvern/apollo/configs/initializer"
	"github.com/gin-gonic/gin"
)

// pong for ping
func pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

// 定义 router
func routerInit() {
	get("ping_pong", "/ping", pong)
	get("home_page", "/", controllers.HomeIndex)
	get("about", "/about", controllers.HomeAboutHandler)

	// captcha
	get("captcha_get", initializer.Captcha.URLPrefix+":id", controllers.GetCaptcha)

	// account
	get("signup", "/signup", controllers.SignupGet)
	post("signup_post", "/signup", controllers.SignUpPost)
	get("signin", "/signin", controllers.SigninGet)
	post("signin_post", "/signin", controllers.SignInPost)
	get("signout", "/signout", controllers.SignOut)

	// user
	get("user_info", "/user/info", pong)

	// url_title
	get("url_title", "/url/title", controllers.QueryTitleFromURL)
	// share
	get("share_new_get", "/share/new", controllers.ShareNewGet)
	post("share_new_post", "/share/new", controllers.ShareNewPost)
}
