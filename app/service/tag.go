package service

import (
	"fmt"
	"sort"
	"strings"

	"github.com/chalvern/apollo/app/model"
	"github.com/chalvern/sugar"
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

// TagsRecommendQuery 检索推荐的标签（目前定义为最新的 30 个标签）
func TagsRecommendQuery() (tags []model.Tag) {
	tags, _, err := TagsQuery(1, 30)
	if err != nil {
		sugar.Warnf("TagsRecommendQuery 出错：%s", err.Error())
	}
	return tags
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

// TagUpdateCount 更新 tag 的数目
// 某个 Tag 下面有多少个分享
func TagUpdateCount(tagName string) error {
	count, err := shareModel.AggregateTagCount(tagName)
	if err != nil {
		return err
	}
	tag, err := tagModel.QueryByName(tagName)
	if err != nil {
		return err
	}
	tagNew := model.Tag{
		Count: count,
	}
	tagNew.ID = tag.ID
	return tagNew.Update()
}

// TagClassifyA 把标签按照一定规则分类排序
func TagClassifyA(sourceTags []model.Tag) []model.Tag {
	unsortedTags := []model.Tag{}
	tagsMap := make(map[string][]model.Tag)
	familyKeys := []string{}
	for _, tag := range sourceTags {
		// 检出第一级标签
		if tag.Hierarchy == 0 && tag.Parent == "" {
			namedTagSlice, ok := tagsMap[tag.Name]
			if !ok {
				namedTagSlice = []model.Tag{}
				familyKeys = append(familyKeys, tag.Name)
			}
			tagsMap[tag.Name] = append(namedTagSlice, tag)
		} else {
			unsortedTags = append(unsortedTags, tag)
		}
	}
	// 把未分类的标签找到合适的 family 进行处理
	for _, tag := range unsortedTags {
		if tag.Parent == "" {
			continue
		}
		family, ok := tagsMap[tag.Parent]
		if ok {
			tagsMap[tag.Parent] = append(family, tag)
		}
	}

	// 归并
	sort.Strings(familyKeys)
	dstTags := []model.Tag{}
	for _, key := range familyKeys {
		for _, tag := range tagsMap[key] {
			dstTags = append(dstTags, tag)
		}
	}
	return dstTags
}
