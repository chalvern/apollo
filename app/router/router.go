package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// pong for ping
func pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
