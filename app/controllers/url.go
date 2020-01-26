package controllers

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/chalvern/apollo/app/service"
	"github.com/chalvern/sugar"
	"github.com/gin-gonic/gin"
)

// QueryTitleFromURL 根据用户传入的 URL 获取对应的标题
func QueryTitleFromURL(c *gin.Context) {
	urlRaw := c.Query("url")
	validURL := govalidator.IsURL(urlRaw)
	if !validURL {
		c.String(http.StatusOK, "URL 不合法，请检查后重试")
		return
	}
	sugar.Infof("QueryTitleFromURL: %s", urlRaw)

	title, err := service.QueryTitleFormURL(urlRaw)
	if err != nil {
		sugar.Errorf("检索标题出错：%s", err.Error())
		c.String(http.StatusOK, "自动获取出错")
		return
	}

	c.String(http.StatusOK, title)
}
