package controllers

import (
	"github.com/chalvern/apollo/configs/initializer"
	"github.com/gin-gonic/gin"
)

// GetCaptcha 获取验证码
func GetCaptcha(c *gin.Context) {
	initializer.Captcha.Handler(c)
	return
}
