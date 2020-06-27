package controllers

import (
	"net/http"
	"strconv"

	"github.com/chalvern/apollo/app/model"
	"github.com/chalvern/apollo/app/service"
	"github.com/chalvern/sugar"
	"github.com/gin-gonic/gin"
)

// ChecklistNewPost 创建一个新的检查项
func ChecklistNewPost(c *gin.Context) {
	c.Set(PageTitle, "新检查项")
	form := struct {
		ShareID uint   `form:"share_id" binding:"required"`
		PreID   uint   `form:"pre_id"`
		Title   string `form:"title" binding:"required,lengte=1,lenlte=200"`
	}{}

	if errs := c.ShouldBind(&form); errs != nil {
		sugar.Warnf("ChecklistNewPost Bind form Error: %s", errs.Error())
		html(c, http.StatusOK, "notify/error.tpl", gin.H{
			"FlashError": "请检查是否填写正确",
		})
		return
	}

	share, err := service.ShareQueryByID(form.ShareID)
	if err != nil {
		sugar.Warnf("ChecklistNewPost 未找到对应的分享（%d） form Error: %s", form.ShareID, err.Error())
		html(c, http.StatusOK, "notify/error.tpl", gin.H{
			"FlashError": "请检查分享的 ID 是否填写正确",
		})
		return
	}

	user := c.MustGet("user").(*model.User)
	checklist := model.Checklist{
		UserID:  user.ID,
		Title:   form.Title,
		ShareID: form.ShareID,
		PrevID:  form.PreID,
	}

	if err := service.ChecklistCreate(&checklist, &share); err != nil {
		sugar.Errorf("ShareNewPost Create Error: %s", err.Error())
		htmlOfOk(c, "notify/error.tpl", gin.H{
			"FlashError": "保存出现了问题，请检查提交内容后稍后重试",
		})
		return
	}

	htmlOfOk(c, "notify/success.tpl", gin.H{
		"Info":         "发布成功!",
		"Timeout":      1,
		"RedirectURL":  "/share/detail?id=" + strconv.Itoa(int(share.ID)),
		"RedirectName": share.Title,
	})

}

// ChecklistUpdate 更新检查项
func ChecklistUpdate(c *gin.Context) {
	c.Set(PageTitle, "更新检查项")
	form := struct {
		ShareID     uint   `form:"share_id" binding:"required"`
		ChecklistID uint   `form:"checklist_id" binding:"required"`
		Title       string `form:"title" binding:"required,lengte=1,lenlte=200"`
	}{}
	if errs := c.ShouldBind(&form); errs != nil {
		sugar.Warnf("ChecklistUpdate Bind form Error: %s", errs.Error())
		html(c, http.StatusOK, "notify/error.tpl", gin.H{
			"FlashError": "请检查是否填写正确",
		})
		return
	}
	user := c.MustGet("user").(*model.User)
	err := service.ChecklistUpdate(form.ChecklistID, form.Title, user)
	if err != nil {
		sugar.Errorf("ChecklistUpdate Update Error: %s", err.Error())
		htmlOfOk(c, "notify/error.tpl", gin.H{
			"FlashError": "更新出现了问题，请检查提交内容后稍后重试",
		})
		return
	}

	htmlOfOk(c, "notify/success.tpl", gin.H{
		"Info":         "发布成功!",
		"Timeout":      1,
		"RedirectURL":  "/share/detail?id=" + strconv.Itoa(int(form.ShareID)),
		"RedirectName": "原帖子",
	})
}
