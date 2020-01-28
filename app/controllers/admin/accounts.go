package admin

import (
	"net/http"

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
