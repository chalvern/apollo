package router

import (
	"net/http"

	"github.com/chalvern/apollo/app/controllers"
	i "github.com/chalvern/apollo/app/interceptors"
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
	get("signup", "/signup", i.UserMustNotExistMiddleware(), controllers.SignupGet)
	post("signup_post", "/signup", i.UserMustNotExistMiddleware(), controllers.SignUpPost)
	get("signin", "/signin", i.UserMustNotExistMiddleware(), controllers.SigninGet)
	post("signin_post", "/signin", i.UserMustNotExistMiddleware(), controllers.SignInPost)
	get("signout", "/signout", controllers.SignOut)

	// user
	get("user_detail", "/user/detail", controllers.UserInfoHandler)

	// tag
	get("tag_detail", "/tag/detail", controllers.TagInfoHandler)

	// url_title
	get("url_title", "/url/title", controllers.QueryTitleFromURL)

	// share
	get("share_detail", "/share/detail", controllers.ShareDetailGet)
	get("share_direct_jump", "/share/redirect", controllers.ShareRedirect)
	get("share_new_get", "/share/new", i.UserMustExistMiddleware(), controllers.ShareNewGet)
	post("share_new_post", "/share/new", i.UserMustExistMiddleware(), controllers.ShareNewPost)

}
