package controllers

import (
	"net/http"

	"github.com/chalvern/apollo/app/model"
	"github.com/chalvern/apollo/app/service"
	"github.com/chalvern/sugar"
	"github.com/gin-gonic/gin"
)

// TagsListHandler 标签列表
func TagsListHandler(c *gin.Context) {
	c.Set(PageTitle, "标签列表")

	page := service.QueryPage(c)
	tags, allPage, err := service.TagsQueryWithContext(c)
	if err != nil {
		sugar.Errorf("TagsListHandler-获取 Shares 出错:%s", err.Error())
		html(c, http.StatusOK, "notify/error.tpl", gin.H{
			"Timeout": 3,
		})
		return
	}

	html(c, http.StatusOK, "tags/list.tpl", gin.H{
		"Tags":        tags,
		"CurrentPage": page,
		"TotalPage":   allPage,
	})
}

// TagInfoHandler 标签信息页面
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

// TagNewGet 创建
func TagNewGet(c *gin.Context) {
	c.Set(PageTitle, "创建标签")
	htmlOfOk(c, "tags/new.tpl", gin.H{})
}

// TagNewPost 创建
func TagNewPost(c *gin.Context) {
	c.Set(PageTitle, "创建标签")
	form := struct {
		Name      string `form:"name" binding:"required"`
		Hierarchy int    `form:"hierarchy"`
		Parent    string `form:"parent"`
		Desc      string `form:"desc"`
	}{}

	if errs := c.ShouldBind(&form); errs != nil {
		sugar.Warnf("ShareNewPost Bind form Error: %s", errs.Error())
		html(c, http.StatusOK, "tags/new.tpl", gin.H{
			"FlashError": "请检查各字段是否填写正确",
		})
		return
	}
	tag := model.Tag{
		Name: form.Name,
	}
	if form.Hierarchy > 0 {
		tag.Hierarchy = form.Hierarchy
	}
	if form.Parent != "" {
		tag.Parent = form.Parent
	}
	if form.Desc != "" {
		tag.Desc = form.Desc
	}

	if err := service.TagCreate(&tag); err != nil {
		sugar.Errorf("TagNewPost Create Error: %s", err.Error())
		htmlOfOk(c, "tags/new.tpl", gin.H{
			"FlashError": "保存出现了问题，请检查提交内容后稍后重试",
		})
		return
	}
	htmlOfOk(c, "notify/success.tpl", gin.H{
		"Info":         "创建成功!",
		"Timeout":      3,
		"RedirectURL":  "/tag/detail?t=" + tag.Name,
		"RedirectName": tag.Name,
	})
}

// TagEditGet 更新
func TagEditGet(c *gin.Context) {
	c.Set(PageTitle, "更新标签")
	tag, err := service.TagQueryByName(c.Query("t"))
	if err != nil || tag.ID == 0 {
		sugar.Warnf("TagEditGet 未检索到对应的标签，name=%v", c.Query("t"))
		html(c, http.StatusOK, "notify/error.tpl", gin.H{
			"Info": "对应条目找不到了，请看看别的吧",
		})
		return
	}

	htmlOfOk(c, "tags/edit.tpl", gin.H{
		"Tag": &tag,
	})
}

// TagEditPost 更新
func TagEditPost(c *gin.Context) {
	c.Set(PageTitle, "更新标签")
	tagOld, err := service.TagQueryByName(c.Query("t"))
	if err != nil || tagOld.ID == 0 {
		sugar.Warnf("TagEditPost 未检索到对应的标签，name=%v", c.Query("t"))
		html(c, http.StatusOK, "notify/error.tpl", gin.H{
			"Info": "对应条目找不到了，请看看别的吧",
		})
		return
	}

	form := struct {
		Name      string `form:"name" binding:"required"`
		Hierarchy int    `form:"hierarchy"`
		Parent    string `form:"parent"`
		Desc      string `form:"desc"`
	}{}

	if errs := c.ShouldBind(&form); errs != nil {
		sugar.Warnf("TagEditPost Bind form Error: %s", errs.Error())
		html(c, http.StatusOK, "tags/edit.tpl", gin.H{
			"FlashError": "请检查各字段是否填写正确",
		})
		return
	}

	tag := model.Tag{
		Name: form.Name,
	}
	if form.Hierarchy > 0 {
		tag.Hierarchy = form.Hierarchy
	}
	if form.Parent != "" {
		tag.Parent = form.Parent
	}
	if form.Desc != "" {
		tag.Desc = form.Desc
	}
	tag.ID = tagOld.ID

	if err := service.TagUpdates(&tag); err != nil {
		sugar.Errorf("TagEditPost Edit Error: %s", err.Error())
		htmlOfOk(c, "tags/edit.tpl", gin.H{
			"FlashError": "保存出现了问题，请检查提交内容后稍后重试",
		})
		return
	}
	htmlOfOk(c, "notify/success.tpl", gin.H{
		"Info":         "更新成功!",
		"Timeout":      3,
		"RedirectURL":  "/tag/detail?t=" + tag.Name,
		"RedirectName": tag.Name,
	})
}
