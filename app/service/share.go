package service

import "github.com/chalvern/apollo/app/model"

import "fmt"

var (
	shareModel = model.Share{}
)

// SharesQuery 检索分享
func SharesQuery(page, pageSize int, args ...interface{}) (tags []model.Share, total int, err error) {
	offset := (page - 1) * pageSize
	return shareModel.QueryBatch(offset, pageSize, args...)
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
	// 只有本人 或者 管理员 可以修改
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
