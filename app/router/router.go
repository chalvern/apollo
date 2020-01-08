package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// pong for ping
func pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

// 定义 router
func router() {
	get("ping_pong", "/ping", pong)
}
