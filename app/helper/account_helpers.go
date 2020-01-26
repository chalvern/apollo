package helper

import (
	"github.com/chalvern/apollo/app/model"
)

// AccountNormalHelper 用户鉴权
func AccountNormalHelper(u *model.User) bool {
	if u == nil {
		return false
	}
	if u.Priority > 0 {
		return true
	}
	return false
}
