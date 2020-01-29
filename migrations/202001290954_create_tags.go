// Package migrations auto generated file
package migrations

import (
	"github.com/chalvern/apollo/app/model"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func init() {
	migrations = append(migrations, &gormigrate.Migration{

		ID: "202001290954",
		Migrate: func(tx *gorm.DB) error {
			// it's a good pratice to copy the struct inside the function,
			// so side effects are prevented if the original struct changes during the time
			type Tag struct {
				model.Model
				Name      string `gorm:"unique_index"` // 标签名称
				Hierarchy int    `gorm:"default:0"`    // 缩进与阶层
				Parent    string `gorm:"index"`        // 父标签
				Desc      string `gorm:"type:text"`    // 描述
				Count     int    `gorm:""`             // 数量
			}
			return tx.AutoMigrate(&Tag{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable("tags").Error
		},
	})
}
