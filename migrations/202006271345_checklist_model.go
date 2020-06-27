// Package migrations auto generated file
package migrations

import (
	"github.com/chalvern/apollo/app/model"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func init() {
	migrations = append(migrations, &gormigrate.Migration{

		ID: "202006271345",
		Migrate: func(tx *gorm.DB) error {
			// it's a good pratice to copy the struct inside the function,
			// so side effects are prevented if the original struct changes during the time
			type Checklist struct {
				model.Model
				ShareID uint   `gorm:"index:idx_share_user_id"` // 分享 ID
				UserID  uint   `gorm:"index:idx_share_user_id"` // 用户 ID
				Title   string `gorm:"type:varchar(256)"`       // 主要描述，一般认为一个 checklist 描述在 200 字内就足够了
				// 维护一个双链表
				PrevID uint `gorm:"default:0"` // 前一个的 checklistID
				PostID uint `gorm:"default:0"` // 后一个的 checklistID
			}
			return tx.AutoMigrate(&Checklist{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable("checklists").Error
		},
	})
}
