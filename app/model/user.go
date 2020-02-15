package model

import (
	"fmt"
	"time"
)

// 用户身份
const (
	UserPriorityUnValid = 0       // 未经任何认证
	UserPriorityCommon  = 1 << 0  // 普通用户
	UserPriorityAdmin   = 1 << 10 // 管理员
	UserPrioritySuper   = 1 << 11 // 超级管理员

	UserPriorityAllValidMask = ^0 // 超级超级用户，全 1
	UserPriorityManagerMask  = UserPriorityAdmin | UserPrioritySuper
)

// User 存放用户信息
type User struct {
	Model
	// Username 用户名，限定字符数目
	Email    string `gorm:"type:varchar(100);unique_index" json:"email"`
	Password string `gorm:"type:varchar(100);" json:"-"`

	Nickname string `gorm:"type:varchar(50)" json:"nickname"` // 昵称

	EmailValidToken string `gorm:"type:varchar(50)" json:"-"` // 邮件认证的token
	EmailVarified   bool   `gorm:"" json:"-"`                 // 邮件已认证

	ResetPasswordKey string     `gorm:"type:varchar(128)" json:"-"` // 重置密码所需要的key
	BannedTime       *time.Time `gorm:""`                           // 被拉黑到什么时候

	// Priority 用户优先级，暂时使用这个字段给用户赋权
	// 比如 普通用户给 2^0，admin 给 2^10
	Priority int `gorm:"default:0" json:"priority"` // 权限优先级
}

// QueryBatch 检索一组
func (u *User) QueryBatch(offset, pageSize int, args ...interface{}) (users []User, total int, err error) {
	db := dbArgs(mydb, args...)
	err = db.Model(User{}).Count(&total).Error
	if err != nil {
		return
	}
	err = db.Offset(offset).
		Limit(pageSize).Order("id desc").
		Find(&users).Error
	return
}

// FindByEmail 根据 Email 检索用户
func (u *User) FindByEmail(email string) (User, error) {
	user := User{}
	err := mydb.Model(u).Where("email=?", email).
		First(&user).Error
	return user, err
}

// FindByUID 通过 ID 检索用户
func (u *User) FindByUID(uid interface{}) (User, error) {
	user := User{}
	err := mydb.Model(u).Where("id=?", uid).
		First(&user).Error
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

// Update 更新
func (u *User) Update() error {
	if u.ID == 0 {
		return fmt.Errorf("用户 更新必须设置 ID")
	}
	return mydb.Model(u).Updates(u).Error
}
