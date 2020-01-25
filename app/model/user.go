package model

import "time"

// User 存放用户信息
type User struct {
	Model
	// Username 用户名，限定字符数目
	Username string `gorm:"type:varchar(30);unique_index" json:"username"`
	Password string `gorm:"type:varchar(100);" json:"-"`
	Email    string `gorm:"type:varchar(100);unique_index" json:"email"`
	Phone    string `gorm:"type:varchar(50);index"` // 电话

	PhoneVarified bool `gorm:"" json:"-"` // 手机号已确认
	EmailVarified bool `gorm:"" json:"-"` // 邮件已认证

	ResetPasswordKey string     `gorm:"type:varchar(128)" json:"-"` // 重置密码所需要的key
	BannedTime       *time.Time `gorm:""`                           // 被拉黑到什么时候
}
