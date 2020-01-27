package controllers

import (
	"net/http"
	"strconv"

	"github.com/chalvern/sugar"
	"github.com/gin-gonic/gin"
)

// 上下文相关的几个常量
const (
	// Apollo 相关
	ApolloCode = "apollo_code"
	ApolloTmpl = "apollo_tmpl"
	ApolloObj  = "apollo_obj"

	PageTitle = "PageTitle"
)

func html(c *gin.Context, code int, tmpl string, obj gin.H) {
	c.Set(ApolloCode, code)
	c.Set(ApolloTmpl, tmpl)
	c.Set(ApolloObj, obj)
}

// htmlOfOk 返回 ok 的内容，包含了用户信息
func htmlOfOk(c *gin.Context, tmpl string, obj gin.H) {
	html(c, http.StatusOK, tmpl, obj)
}

// HTMLOfOK 返回OK内容
func HTMLOfOK(c *gin.Context, tmpl string, obj gin.H) {
	htmlOfOk(c, tmpl, obj)
}

// PageNotFound 页面不存在（404）
func PageNotFound(c *gin.Context) {
	c.Set(PageTitle, "页面不存在")
	sugar.Infof("PageNotFound:%s", c.Request.URL.Path)
	htmlOfOk(c, "notify/not_found.tpl", gin.H{
		"FlashError": "未找到对应的页面",
	})
	return
}

// query 参数

// queryPage 抽取 page 数目
// 默认从 1 开始计数
func queryPage(c *gin.Context) int {
	pageString := c.Query("page")
	page := 1
	if pageString != "" {
		p, err := strconv.Atoi(pageString)
		if err == nil {
			page = p
		}
	}
	return page
}

func queryPageSize(c *gin.Context) int {
	pageSizeString := c.Query("page_size")
	pageSize := 20
	if pageSizeString != "" {
		p, err := strconv.Atoi(pageSizeString)
		if err == nil {
			pageSize = p
		}
	}
	return pageSize
}
