package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeIndex Home 首页
func HomeIndex(c *gin.Context) {
	c.Set(PageTitle, "见周边（正见生活、品鉴生活，推荐生活）")
	tabString := c.Query("t")
	if tabString == "" {
		tabString = "0"
	}

	html(c, http.StatusOK, "home/index.tpl", gin.H{})
}
