// Package migrations auto generated file
package migrations

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func init() {
	migrations = append(migrations, &gormigrate.Migration{

		ID: "202006271149",
		Migrate: func(tx *gorm.DB) error {
			// it's a good pratice to copy the struct inside the function,
			// so side effects are prevented if the original struct changes during the time
			type Share struct {
				ChecklistID uint `gorm:"default:0"` // 检查项头
			}
			return tx.AutoMigrate(&Share{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Table("shares").DropColumn("checklist_id").Error
		},
	})
}
