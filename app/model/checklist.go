package model

import (
	"fmt"
)

// Checklist 检查项
type Checklist struct {
	Model
	ShareID uint   `gorm:"index:idx_share_user_id"` // 分享 ID
	UserID  uint   `gorm:"index:idx_share_user_id"` // 用户 ID
	Title   string `gorm:"type:varchar(256)"`       // 主要描述，一般认为一个 checklist 描述在 200 字内就足够了
	// 维护一个双链表
	PrevID uint `gorm:"default:0"` // 前一个的 checklistID
	PostID uint `gorm:"default:0"` // 后一个的 checklistID
}

// QueryBatch 检索一组
func (c *Checklist) QueryBatch(args ...interface{}) (checklists []*Checklist, err error) {
	db := dbArgs(mydb, args...)
	if err != nil {
		return
	}
	err = db.Model(Checklist{}).Find(&checklists).Error
	return
}

// Create 创建
func (c *Checklist) Create() error {
	return mydb.Save(c).Error
}

// Updates 更新
func (c *Checklist) Updates(values interface{}) error {
	if c.ID == 0 {
		return fmt.Errorf("更新必须设置 ID")
	}
	return mydb.Model(c).Updates(values).Error
}

// Update 更新单个字段
func (c *Checklist) Update(column string, value interface{}) error {
	if c.ID == 0 {
		return fmt.Errorf("更新必须设置 ID")
	}
	return mydb.Model(c).Update(column, value).Error
}
