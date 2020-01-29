package interceptors

import (
	"net/http"

	"github.com/chalvern/apollo/app/model"
	"github.com/gin-gonic/gin"
)

// UserPriorityMiddleware 检查用户登录
func UserPriorityMiddleware(pMast int) gin.HandlerFunc {
	return func(c *gin.Context) {
		accountRaw, exists := c.Get("user")
		if !exists {
			c.HTML(http.StatusOK, "notify/error.tpl", gin.H{
				"PageTitle":  "出错了",
				"FlashError": "用户未登录",
			})
			c.Abort()
			return
		}
		account := accountRaw.(*model.User)
		if account.Priority&pMast == 0 {
			c.HTML(http.StatusOK, "notify/error.tpl", gin.H{
				"PageTitle":  "出错了",
				"FlashError": "用户非管理员",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
