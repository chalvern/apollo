package service

import (
	"fmt"
	"strings"

	"github.com/chalvern/apollo/app/model"
	"github.com/gin-gonic/gin"
)

var (
	commentModel = model.Comment{}
)

// CommentsQueryWithContext 根据 url query 参数检索
func CommentsQueryWithContext(c *gin.Context, preloadUser bool, args ...interface{}) (comments []model.Comment, allPage int, err error) {

	page := queryPage(c)
	pageSize := queryPageSize(c)

	argS, argArray := argsInit(args...)
	if statusStr := c.Query("s"); statusStr != "" {
		argS = append(argS, "status=?")
		argArray = append(argArray, statusStr)
	}

	argArray[0] = strings.Join(argS, "AND")

	offset := (page - 1) * pageSize
	comments, total, err := commentModel.QueryBatch(offset, pageSize, preloadUser, args...)
	allPage = total/pageSize + 1
	return
}

// CommentCreate 创建
func CommentCreate(comment *model.Comment) error {
	if comment.UserID == 0 || comment.ShareID == 0 {
		return fmt.Errorf("创建分享必须设置用户 ID 和 分享 ID")
	}
	return comment.Create()
}

// CommentUpdates 更新
func CommentUpdates(comment *model.Comment, user *model.User) error {
	if comment.ID == 0 {
		return fmt.Errorf("更新必须是已存在的分享内容")
	}
	// 只能本人或者管理员有修改分享的权限
	if comment.UserID == user.ID || user.Priority&model.UserPriorityManager != 0 {
		return comment.Update()
	}

	return fmt.Errorf("越权！用户 %s 试图修改分享 %d", user.Email, comment.ID)
}
