// Package migrations auto generated file
package migrations

import (
	"github.com/chalvern/apollo/app/model"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func init() {
	migrations = append(migrations, &gormigrate.Migration{

		ID: "202001262309",
		Migrate: func(tx *gorm.DB) error {
			// it's a good pratice to copy the struct inside the function,
			// so side effects are prevented if the original struct changes during the time
			type Share struct {
				model.Model
				UserID     uint   `gorm:"index"`             // 用户 ID
				URL        string `gorm:"varchar(1024)"`     // URL
				Title      string `gorm:"varchar(100)"`      // 分享的文章标题
				Review     string `gorm:"type:text"`         // 评论
				Status     int    `gorm:"default:0"`         // 状态
				Tag        string `gorm:"index;varchar(30)"` // 标签
				ClickCount int    `gorm:"default:0"`         // 点击数量(浏览数量)
				StarCount  int    `gorm:"default:0"`         // 赞的数量
			}
			return tx.AutoMigrate(&Share{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable("shares").Error
		},
	})
}
