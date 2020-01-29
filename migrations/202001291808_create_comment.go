// Package migrations auto generated file
package migrations

import (
	"github.com/chalvern/apollo/app/model"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func init() {
	migrations = append(migrations, &gormigrate.Migration{

		ID: "202001291808",
		Migrate: func(tx *gorm.DB) error {
			// it's a good pratice to copy the struct inside the function,
			// so side effects are prevented if the original struct changes during the time
			type Comment struct {
				model.Model
				ShareID uint   `gorm:"index:idx_share_user_id"` // 分享 ID
				UserID  uint   `gorm:"index:idx_share_user_id"` // 用户 ID
				Reply   string `gorm:"type:text"`               // 回复细节
				Number  int    `gorm:""`                        // 楼层高度
			}
			return tx.AutoMigrate(&Comment{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable("comments").Error
		},
	})
}
