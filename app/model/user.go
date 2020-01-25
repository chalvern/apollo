package model

import "time"

// User 存放用户信息
type User struct {
	Model
	// Username 用户名，限定字符数目
	Email    string `gorm:"type:varchar(100);unique_index" json:"email"`
	Password string `gorm:"type:varchar(100);" json:"-"`

	EmailVarified bool `gorm:"" json:"-"` // 邮件已认证

	ResetPasswordKey string     `gorm:"type:varchar(128)" json:"-"` // 重置密码所需要的key
	BannedTime       *time.Time `gorm:""`                           // 被拉黑到什么时候
}
