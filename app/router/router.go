package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// pong for ping
func pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

// Init initialize engine
func Init(r *gin.Engine) *gin.Engine {
	r.GET("/ping", pong)
	return r
}
