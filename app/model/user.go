package model

import (
	"fmt"
	"time"
)

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

// FindByEmail 根据 Email 检索用户
func (u *User) FindByEmail(email string) (User, error) {
	user := User{}
	err := mydb.Model(u).First(&user).Error
	return user, err
}

// Create 创建用户
func (u *User) Create() error {
	if len(u.Email) <= 0 || len(u.Email) > 100 {
		return fmt.Errorf("传入的邮箱 %s 长度非法", u.Email)
	}
	if u.Password == "" {
		return fmt.Errorf("创建用户时密码不能为空")
	}
	return mydb.Save(u).Error
}
