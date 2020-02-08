package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeIndex 首页
func HomeIndex(c *gin.Context) {
	c.Set(PageTitle, "管理员首页")

	html(c, http.StatusOK, "admin/index.tpl", gin.H{})
}
