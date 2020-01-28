package service

import (
	"fmt"
	"strings"

	"github.com/chalvern/apollo/app/model"
	"github.com/gin-gonic/gin"
)

var (
	shareModel = model.Share{}
)

// SharesQueryWithContext 根据 url query 参数检索
func SharesQueryWithContext(c *gin.Context, preloadUser bool, args ...interface{}) (shares []model.Share, allPage int, err error) {

	page := queryPage(c)
	pageSize := queryPageSize(c)

	argS, argArray := argsInit(args...)
	if statusStr := c.Query("s"); statusStr != "" {
		argS = append(argS, "status=?")
		argArray = append(argArray, statusStr)
	}

	argArray[0] = strings.Join(argS, "AND")

	shares, total, err := SharesQuery(page, pageSize, preloadUser, argArray...)
	allPage = total/pageSize + 1
	return
}

// ShareQueryByID 根据 ID 检索对应的分享
func ShareQueryByID(id interface{}) (model.Share, error) {
	return shareModel.QueryByID(id)
}

// SharesQuery 检索分享
func SharesQuery(page, pageSize int, preloadUser bool, args ...interface{}) (tags []model.Share, total int, err error) {
	offset := (page - 1) * pageSize
	return shareModel.QueryBatch(offset, pageSize, preloadUser, args...)
}

// ShareCreate 创建分享
func ShareCreate(share *model.Share) error {
	if share.Tag == "" {
		return fmt.Errorf("分享(标题 %s)必须设定标签", share.Title)
	}
	if share.UserID == 0 {
		return fmt.Errorf("创建分享必须设置用户 ID")
	}
	return share.Create()
}

// ShareUpdates 更新分享
func ShareUpdates(share *model.Share, user *model.User) error {
	if share.ID == 0 {
		return fmt.Errorf("更新分享必须是已存在的分享内容")
	}
	// 只能本人或者管理员有修改分享的权限
	if share.UserID == user.ID || user.Priority&model.UserPriorityManager != 0 {
		return share.Update()
	}

	return fmt.Errorf("越权！用户 %s 试图修改分享 %d", user.Email, share.ID)
}

// ShareClicked 被点击一次
func ShareClicked(shareID uint) error {
	return shareModel.Click(shareID)
}

// ShareStared 被点赞
func ShareStared(shareID uint) error {
	return shareModel.Star(shareID)
}
