package service

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// 初始化检索的参数
func argsInit(args ...interface{}) ([]string, []interface{}) {
	argStringArray := []string{}
	argArray := []interface{}{
		"",
	}
	if len(args) > 1 {
		argStringArray = append(argStringArray, (args[0].([]string))...)
		argArray = append(argArray, args[1:]...)
	}
	return argStringArray, argArray
}

// query 参数

// queryPage 抽取 page 数目
// 默认从 1 开始计数
func queryPage(c *gin.Context) int {
	pageString := c.Query("page")
	page := 1
	if pageString != "" {
		p, err := strconv.Atoi(pageString)
		if err == nil {
			page = p
		}
	}
	return page
}

// QueryPage 从 gin.Context 中提取页码
func QueryPage(c *gin.Context) int {
	return queryPage(c)
}

func queryPageSize(c *gin.Context) int {
	pageSizeString := c.Query("page_size")
	pageSize := 20
	if pageSizeString != "" {
		p, err := strconv.Atoi(pageSizeString)
		if err == nil {
			pageSize = p
		}
	}
	return pageSize
}

// QueryPageSize 从 gin.Context 中提取每页数量
func QueryPageSize(c *gin.Context) int {
	return queryPageSize(c)
}
