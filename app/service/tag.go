package service

import (
	"fmt"
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

// TagQueryByName 根据名称检索
func TagQueryByName(tagName string) (tag model.Tag, err error) {
	return tagModel.QueryByName(tagName)
}

// TagCreate 创建标签
func TagCreate(tag *model.Tag) error {
	tag.Name = strings.ToLower(tag.Name)
	return tag.Create()
}

// TagUpdates 更新分享
func TagUpdates(tag *model.Tag) error {
	if tag.ID == 0 {
		return fmt.Errorf("更新标签必须是已存在的分享内容")
	}
	return tag.Update()
}
