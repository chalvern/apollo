package interceptors

import (
	"net/http"

	"github.com/chalvern/apollo/app/controllers"
	"github.com/chalvern/apollo/app/service"

	"github.com/chalvern/apollo/tools/jwt"

	"github.com/chalvern/sugar"
	"github.com/gin-gonic/gin"
)

// JwtMiddleware 通过 jwt 鉴权
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie(controllers.CookieTag)
		// 如果没有在 cookie 中拿到 token，继续向下执行即可
		if err != nil || len(accessToken) == 0 {
			c.Next()
			return
		}
		claim, err := jwt.ParseToken(accessToken)
		if err != nil {
			sugar.Warnf("AuthOfJwt 鉴权失败(AuthOfJwt)：%s", err.Error())
			c.Abort()
			return
		}
		emailString := claim["email"].(string)

		u, err := service.UserFindByEmail(emailString)
		if err != nil {
			sugar.Warnf("UserFindByEmail (email= %s )失败：%s", emailString, err.Error())
			c.Next()
			return
		}
		c.Set("user", u)
		c.Next()
	}
}

// UserMustExistMiddleware 检查用户登录
func UserMustExistMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, exists := c.Get("user"); !exists {
			c.HTML(http.StatusOK, "notify/error.tpl", gin.H{
				"PageTitle":  "出错了",
				"FlashError": "用户未登录",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

// UserMustNotExistMiddleware 检查用户未登录
func UserMustNotExistMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if account, exists := c.Get("user"); exists {
			c.HTML(http.StatusOK, "notify/error.tpl", gin.H{
				"PageTitle":  "出错了",
				"FlashError": "用户已经登录",
				"Account":    account,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
