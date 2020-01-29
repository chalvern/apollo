package controllers

import (
	"net/http"
	"strconv"

	"github.com/chalvern/apollo/app/model"
	"github.com/chalvern/apollo/app/service"
	"github.com/chalvern/sugar"
	"github.com/gin-gonic/gin"
)

// CommentNewPost 创建新的评论
func CommentNewPost(c *gin.Context) {
	c.Set(PageTitle, "新评论")
	form := struct {
		ShareID uint   `form:"share_id" binding:"required"`
		Replay  string `form:"replay" binding:"required,lengte=1,lenlte=800"`
	}{}

	if errs := c.ShouldBind(&form); errs != nil {
		sugar.Warnf("CommentNewPost Bind form Error: %s", errs.Error())
		html(c, http.StatusOK, "notify/error.tpl", gin.H{
			"FlashError": "请检查是否填写正确",
		})
		return
	}

	share, err := service.ShareQueryByID(form.ShareID)
	if err != nil {
		sugar.Warnf("CommentNewPost 未找到对应的分享（%d） form Error: %s", form.ShareID, err.Error())
		html(c, http.StatusOK, "notify/error.tpl", gin.H{
			"FlashError": "请检查分享的 ID 是否填写正确",
		})
		return
	}

	user := c.MustGet("user").(*model.User)
	comment := model.Comment{
		UserID:  user.ID,
		Reply:   form.Replay,
		ShareID: form.ShareID,
		Number:  share.CommentCount,
	}

	if err := service.CommentCreate(&comment); err != nil {
		sugar.Errorf("ShareNewPost Create Error: %s", err.Error())
		htmlOfOk(c, "notify/error.tpl", gin.H{
			"FlashError": "保存出现了问题，请检查提交内容后稍后重试",
		})
		return
	}

	err = service.ShareComment(share.ID)
	if err != nil {
		sugar.Errorf("记录回复数量出错 %d", share.ID)
	}

	htmlOfOk(c, "notify/success.tpl", gin.H{
		"Info":         "发布成功!",
		"Timeout":      1,
		"RedirectURL":  "/share/detail?id=" + strconv.Itoa(int(share.ID)),
		"RedirectName": share.Title,
	})

}
