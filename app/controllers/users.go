package controllers

import (
	"github.com/chalvern/apollo/app/service"
	"github.com/chalvern/sugar"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserInfoHandler 用户信息页面
func UserInfoHandler(c *gin.Context) {
	c.Set(PageTitle, "用户信息")
	uidString := c.Query("uid")
	user, err := service.UserFindByUID(uidString)
	if err != nil {
		sugar.Warnf("UserInfo-获取用户出错:%s", err.Error())
		html(c, http.StatusOK, "notify/error.tpl", gin.H{
			"FlashError": "未找到对应的用户",
		})
		return
	}

	page := queryPage(c)
	pageSize := queryPageSize(c)
	shares, allPage, err := service.SharesQueryWithContext(c, page, pageSize, true, user.ID)
	if err != nil {
		sugar.Errorf("UserInfo-获取 Shares 出错:%s", err.Error())
		return
	}
	htmlOfOk(c, "users/info.tpl", gin.H{
		"Shares":      shares,
		"User":        user,
		"CurrentPage": page,
		"TotalPage":   allPage,
	})
}
