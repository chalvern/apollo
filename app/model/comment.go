package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Comment 评论
type Comment struct {
	Model
	ShareID uint   `gorm:"index:idx_share_user_id"` // 分享 ID
	UserID  uint   `gorm:"index:idx_share_user_id"` // 用户 ID
	User    *User  `gorm:"foreignkey:UserID"`       // 用户
	Reply   string `gorm:"type:text"`               // 回复细节
	Number  int    `gorm:""`                        // 楼层高度
}

// QueryBatch 检索一组
func (c *Comment) QueryBatch(offset, pageSize int, userPreload bool, args ...interface{}) (comments []Comment, total int, err error) {
	db := dbArgs(mydb, args...)
	err = db.Model(Comment{}).Count(&total).Error
	if err != nil {
		return
	}
	if userPreload {
		db = db.Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id,nickname")
		})
	}
	err = db.Model(Comment{}).Offset(offset).
		Limit(pageSize).Order("id asc").
		Find(&comments).Error
	return
}

// Create 创建
func (c *Comment) Create() error {
	return mydb.Save(c).Error
}

// Update 更新
func (c *Comment) Update() error {
	if c.ID == 0 {
		return fmt.Errorf("更新必须设置 ID")
	}
	return mydb.Model(c).Updates(c).Error
}
