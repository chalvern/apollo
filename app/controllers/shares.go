package controllers

import (
	"net/http"
	"strconv"

	"github.com/chalvern/apollo/app/helper"
	"github.com/chalvern/apollo/app/model"
	"github.com/chalvern/apollo/app/service"
	"github.com/chalvern/sugar"
	"github.com/gin-gonic/gin"
)

// ShareDetailGet 点击一个分享
func ShareDetailGet(c *gin.Context) {
	c.Set(PageTitle, "分享详情")
	share, err := service.ShareQueryByID(c.Query("id"))
	if err != nil || share.ID == 0 {
		sugar.Warnf("ShareDetailGet 未检索到对应的分享，id=%v", c.Query("id"))
		html(c, http.StatusOK, "notify/error.tpl", gin.H{
			"Info": "对应条目找不到了，请看看别的吧",
		})
		return
	}
	// 非用户本身才加 1
	userTmp, exists := c.Get("user")
	if exists {
		user := userTmp.(*model.User)
		if share.UserID != user.ID {
			share.Click(share.ID)
		}
	}
	htmlOfOk(c, "shares/detail.tpl", gin.H{
		"Share": &share,
	})
}

// ShareRedirect 直接跳转
func ShareRedirect(c *gin.Context) {
	share, err := service.ShareQueryByID(c.Query("id"))
	if err != nil || share.ID == 0 {
		sugar.Warnf("ShareRedirect 未检索到对应的分享，id=%v", c.Query("id"))
		html(c, http.StatusOK, "notify/error.tpl", gin.H{
			"Info": "对应条目找不到了，请看看别的吧",
		})
		return
	}
	// 非用户本身才加 1
	userTmp, exists := c.Get("user")
	if exists {
		user := userTmp.(*model.User)
		if share.UserID != user.ID {
			share.Click(share.ID)
		}
	}
	c.Redirect(http.StatusSeeOther, share.URL)
}

// ShareNewGet 创建分享的表单
func ShareNewGet(c *gin.Context) {
	c.Set(PageTitle, "创建分享")
	htmlOfOk(c, "shares/new.tpl", gin.H{})
}

// ShareNewPost 创建分享
func ShareNewPost(c *gin.Context) {
	c.Set(PageTitle, "创建分享")
	form := struct {
		URL    string `form:"url" binding:"required,url"`
		Title  string `form:"title" binding:"required,lengte=1"`
		Review string `form:"review" binding:"required,lenlte=800"`
		Tag    string `form:"tag" binding:"required"`
	}{}

	if errs := c.ShouldBind(&form); errs != nil {
		sugar.Warnf("ShareNewPost Bind form Error: %s", errs.Error())
		html(c, http.StatusOK, "shares/new.tpl", gin.H{
			"FlashError": "请检查URL、标题、评论等内容及格式是否填写正确",
		})
		return
	}

	user := c.MustGet("user").(*model.User)
	share := model.Share{
		UserID: user.ID,
		URL:    form.URL,
		Title:  form.Title,
		Review: form.Review,
		Tag:    form.Tag,
	}

	if err := service.ShareCreate(&share); err != nil {
		sugar.Errorf("ShareNewPost Create Error: %s", err.Error())
		htmlOfOk(c, "shares/new.tpl", gin.H{
			"FlashError": "保存出现了问题，请检查提交内容后稍后重试",
		})
		return
	}
	htmlOfOk(c, "notify/success.tpl", gin.H{
		"Info":         "发布成功!",
		"Timeout":      3,
		"RedirectURL":  "/share/detail?id=" + strconv.Itoa(int(share.ID)),
		"RedirectName": share.Title,
	})
}

// ShareEditGet 创建分享的表单
func ShareEditGet(c *gin.Context) {
	c.Set(PageTitle, "编辑分享")
	share, err := service.ShareQueryByID(c.Query("id"))
	if err != nil || share.ID == 0 {
		sugar.Warnf("ShareDetailGet 未检索到对应的分享，id=%v", c.Query("id"))
		html(c, http.StatusOK, "notify/error.tpl", gin.H{
			"Info": "对应条目找不到了，请看看别的吧",
		})
		return
	}
	user := c.MustGet("user").(*model.User)
	// 只能本人或者管理员有修改分享的权限
	if user.ID != share.UserID || !helper.AccountManagerHelper(user) {
		sugar.Warnf("ShareEditPost 用户 %d 没有修改 id=%v 的权限", user.ID, c.Query("id"))
		html(c, http.StatusOK, "notify/error.tpl", gin.H{
			"Info": "没有权限修改此分享内容",
		})
		return
	}

	htmlOfOk(c, "shares/edit.tpl", gin.H{
		"Share": &share,
	})
}

// ShareEditPost 更新分享
func ShareEditPost(c *gin.Context) {
	c.Set(PageTitle, "更新分享")

	shareOld, err := service.ShareQueryByID(c.Query("id"))
	if err != nil || shareOld.ID == 0 {
		sugar.Warnf("ShareDetailGet 未检索到对应的分享，id=%v", c.Query("id"))
		html(c, http.StatusOK, "notify/error.tpl", gin.H{
			"Info": "对应条目找不到了，请看看别的吧",
		})
		return
	}
	user := c.MustGet("user").(*model.User)
	// 只能本人或者管理员有修改分享的权限
	if user.ID != shareOld.UserID || !helper.AccountManagerHelper(user) {
		sugar.Warnf("ShareEditPost 用户 %d 没有修改 id=%v 的权限", user.ID, c.Query("id"))
		html(c, http.StatusOK, "notify/error.tpl", gin.H{
			"Info": "没有权限修改此分享内容",
		})
		return
	}

	form := struct {
		URL    string `form:"url" binding:"required,url"`
		Title  string `form:"title" binding:"required,lengte=1"`
		Review string `form:"review" binding:"required,lenlte=800"`
		Tag    string `form:"tag" binding:"required"`
	}{}
	if errs := c.ShouldBind(&form); errs != nil {
		sugar.Warnf("ShareNewPost Bind form Error: %s", errs.Error())
		html(c, http.StatusOK, "shares/edit.tpl", gin.H{
			"FlashError": "请检查URL、标题、评论等内容及格式是否填写正确",
		})
		return
	}

	share := model.Share{}
	share.ID = shareOld.ID
	share.URL = form.URL
	share.Title = form.Title
	share.Review = form.Review
	share.Tag = form.Tag

	if err = service.ShareUpdates(&share, user); err != nil {
		sugar.Errorf("ShareEditPost Update Error: %s", err.Error())
		htmlOfOk(c, "shares/edit.tpl", gin.H{
			"FlashError": "更新出现了问题，请检查提交内容后稍后重试",
		})
		return
	}
	htmlOfOk(c, "notify/success.tpl", gin.H{
		"Info":         "更新成功!",
		"Timeout":      3,
		"RedirectURL":  "/share/detail?id=" + strconv.Itoa(int(share.ID)),
		"RedirectName": share.Title,
	})
}
