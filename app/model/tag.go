package model

import (
	"fmt"
	"strings"
)

// Tag 标签
type Tag struct {
	Model
	Name      string `gorm:"unique_index"` // 标签名称
	Hierarchy int    `gorm:"default:0"`    // 缩进与阶层
	Parent    string `gorm:"index"`        // 父标签
	Desc      string `gorm:"type:text"`    // 描述
	Count     int    `gorm:"default:0"`    // 数量
}

// QueryBatch 检索一页标签
func (t *Tag) QueryBatch(offset, pageSize int, args ...interface{}) (tags []Tag, total int, err error) {
	db := dbArgs(mydb, args...)
	err = db.Model(Tag{}).Count(&total).Error
	if err != nil {
		return
	}
	err = db.Offset(offset).
		Limit(pageSize).Order("id desc").
		Find(&tags).Error
	return
}

// QueryByName 根据名称获取 标签
func (t *Tag) QueryByName(name string) (Tag, error) {
	tag := Tag{}
	err := mydb.Model(tag).Where("name=?", name).First(&tag).Error
	return tag, err
}

// Create 创建标签
func (t *Tag) Create() error {
	if len(strings.TrimSpace(t.Name)) == 0 {
		return fmt.Errorf("标签不能为空字符串")
	}
	return mydb.Save(t).Error
}

// Update 更新
func (t *Tag) Update() error {
	if t.ID == 0 {
		return fmt.Errorf("Tag 更新必须设置 ID")
	}
	if len(strings.TrimSpace(t.Name)) == 0 {
		return fmt.Errorf("标签不能为空字符串")
	}
	return mydb.Model(t).Updates(t).Error
}
