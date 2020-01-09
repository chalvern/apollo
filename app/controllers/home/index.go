package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index Home 首页
func Index(c *gin.Context) {
	pageTitle := "荐周边（正见生活、品鉴生活，推荐生活）"
	tabString := c.Query("t")
	if tabString == "" {
		tabString = "0"
	}

	c.HTML(http.StatusOK, "home/index.tpl", gin.H{
		"title": pageTitle,
	})
}
