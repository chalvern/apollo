package interceptors

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// AfterRouterMiddleware 最后一步渲染
func AfterRouterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		rawCode, ok := c.Get("apollo_code")
		if !ok {
			rawCode = http.StatusOK
		}
		code := rawCode.(int)

		rawTmpl, ok := c.Get("apollo_tmpl")
		if !ok {
			rawTmpl = "home/index.tpl"
		}
		tmpl := rawTmpl.(string)

		rawObj, ok := c.Get("apollo_obj")
		if !ok {
			rawObj = gin.H{}
		}
		obj := rawObj.(gin.H)

		// 页面标题
		pageTitle, ok := c.Get("PageTitle")
		if ok {
			obj["PageTitle"] = pageTitle
		}

		// 用户
		account, ok := c.Get("user")
		if ok {
			obj["Account"] = account
		}
		c.HTML(code, tmpl, obj)
		c.Next()
	}
}
