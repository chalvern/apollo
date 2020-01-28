package helper

import (
	"github.com/chalvern/apollo/app/model"
)

// AccountNormalHelper 用户鉴权
// 用户拥有基本的权限（用户初始化后 priority=0，不具有基本权限）
func AccountNormalHelper(u *model.User) bool {
	if u == nil {
		return false
	}
	if u.Priority > 0 {
		return true
	}
	return false
}

// AccountManagerHelper 用户具有 管理者（manager）权限
func AccountManagerHelper(u *model.User) bool {
	if u == nil {
		return false
	}
	if u.Priority&model.UserPriorityManager > 0 {
		return true
	}
	return false
}

// AccountHasShareEditAuthority 用户对分享是否有编辑权限
func AccountHasShareEditAuthority(s *model.Share, u *model.User) bool {
	if u == nil || s == nil {
		return false
	}
	if s.UserID == u.ID || u.Priority&model.UserPriorityManager > 0 {
		return true
	}
	return false
}
