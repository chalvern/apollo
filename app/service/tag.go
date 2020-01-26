package service

import "github.com/chalvern/apollo/app/model"

import "strings"

var (
	tagModel = model.Tag{}
)

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
