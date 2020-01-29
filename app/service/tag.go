package service

import (
	"strings"

	"github.com/chalvern/apollo/app/model"
	"github.com/gin-gonic/gin"
)

var (
	tagModel = model.Tag{}
)

// TagsQueryWithContext 根据 url query 参数检索
func TagsQueryWithContext(c *gin.Context, args ...interface{}) (tags []model.Tag, allPage int, err error) {

	page := queryPage(c)
	pageSize := queryPageSize(c)

	argS, argArray := argsInit(args...)
	argArray[0] = strings.Join(argS, "AND")

	tags, total, err := TagsQuery(page, pageSize, argArray...)
	allPage = total/pageSize + 1
	return
}

// TagsQuery 检索标签
func TagsQuery(page, pageSize int, args ...interface{}) (tags []model.Tag, total int, err error) {
	offset := (page - 1) * pageSize
	return tagModel.QueryBatch(offset, pageSize, args...)
}

// TagCreate 创建标签
func TagCreate(name string) error {
	tag := model.Tag{
		Name: strings.ToLower(name),
	}
	return tag.Create()
}
