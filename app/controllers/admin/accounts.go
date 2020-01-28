package admin

import (
	"net/http"

	"github.com/chalvern/apollo/app/model"
	"github.com/chalvern/apollo/app/service"
	"github.com/chalvern/sugar"
	"github.com/gin-gonic/gin"
)

// AccountsList 账户列表
func AccountsList(c *gin.Context) {
	c.Set(PageTitle, "见周边（正见生活、品鉴生活，推荐生活）")
	page := service.QueryPage(c)
	users, allPage, err := service.UsersQueryWithContext(c)

	if err != nil {
		sugar.Errorf("HomeIndex-获取 Shares 出错:%s", err.Error())
		html(c, http.StatusOK, "notify/error.tpl", gin.H{
			"Timeout": 3,
		})
		return
	}

	html(c, http.StatusOK, "admin/users/list.tpl", gin.H{
		"Users":       users,
		"CurrentPage": page,
		"TotalPage":   allPage,
	})
}

// AccountsEditGet 获取编辑
func AccountsEditGet(c *gin.Context) {
	uidString := c.Query("uid")
	user, _ := service.UserFindByUID(uidString)
	c.Set(PageTitle, user.Nickname)
	html(c, http.StatusOK, "admin/users/edit.tpl", gin.H{
		"User": user,
	})
}

// AccountsEditPost 编辑
func AccountsEditPost(c *gin.Context) {
	uidString := c.Query("uid")
	user, _ := service.UserFindByUID(uidString)
	c.Set(PageTitle, user.Nickname)

	form := struct {
		Priority int `form:"priority" binding:"required"`
	}{}
	c.ShouldBind(&form)

	userNew := model.User{
		Priority: form.Priority,
	}
	userNew.ID = user.ID
	err := service.UserUpdates(&userNew)
	if err != nil {
		sugar.Errorf("更新失败 %v", err)
	}
	html(c, http.StatusOK, "notify/success.tpl", gin.H{
		"Info":         "已更新",
		"Timeout":      3,
		"RedirectURL":  "/admin/account/list",
		"RedirectName": "用户列表",
	})
}
