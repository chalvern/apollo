package model

import "fmt"

// Tag 标签
type Tag struct {
	Model
	Name     string `gorm:"unique_index"` // 标签名称
	ParentID uint   `gorm:"index"`        // 父节点 ID
	Count    int    `gorm:""`             // 数量
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
	if len(t.Name) == 0 {
		return fmt.Errorf("标签不能为空字符串")
	}
	return mydb.Save(t).Error
}
