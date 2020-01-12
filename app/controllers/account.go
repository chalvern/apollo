package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SigninGet 获取登录页面
func SigninGet(c *gin.Context) {
	c.HTML(http.StatusOK, "account/signin.tpl", gin.H{
		"PageTitle": "登录",
	})
}
