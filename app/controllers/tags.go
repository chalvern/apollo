package controllers

import (
	"github.com/chalvern/apollo/app/service"
	"github.com/chalvern/sugar"
	"github.com/gin-gonic/gin"
)

// TagInfoHandler 用户信息页面
func TagInfoHandler(c *gin.Context) {
	tagName := c.Query("t")
	if tagName == "" {
		c.Set(PageTitle, "未指定")
	} else {
		c.Set(PageTitle, tagName)
	}

	page := service.QueryPage(c)

	argS, argArray := argsInit()
	argS = append(argS, "tag=?")
	argArray = append(argArray, tagName)
	argArray[0] = argS
	shares, allPage, err := service.SharesQueryWithContext(c, true, argArray...)
	if err != nil {
		sugar.Errorf("UserInfo-获取 Shares 出错:%s", err.Error())
		return
	}
	htmlOfOk(c, "tags/info.tpl", gin.H{
		"CurrentTag":  tagName,
		"Shares":      shares,
		"CurrentPage": page,
		"TotalPage":   allPage,
	})
}
