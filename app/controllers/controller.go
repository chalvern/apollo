package controllers

import (
	"net/http"

	"github.com/chalvern/sugar"
	"github.com/gin-gonic/gin"
)

// htmlOfOk 返回 ok 的内容，包含了用户信息
func htmlOfOk(c *gin.Context, tmpl, pageTitle string, obj gin.H) {
	account, isExists := c.Get("user")
	if isExists {
		obj["Account"] = account
	}

	obj["PageTitle"] = pageTitle

	c.HTML(http.StatusOK, tmpl, obj)
}

// HTMLOfOK 返回OK内容
func HTMLOfOK(c *gin.Context, tmpl, pageTitle string, obj gin.H) {
	htmlOfOk(c, tmpl, pageTitle, obj)
}

// PageNotFound 页面不存在（404）
func PageNotFound(c *gin.Context) {
	sugar.Infof("PageNotFound:%s", c.Request.URL.Path)
	htmlOfOk(c, "notify/not_found.tpl", "页面不存在", gin.H{
		"FlashError": "未找到对应的页面",
	})
	return
}
