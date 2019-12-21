// Package migrations auto generated file
package migrations

import (
	"time"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func init() {
	migrations = append(migrations, &gormigrate.Migration{

		ID: "201912212307",
		Migrate: func(tx *gorm.DB) error {
			// it's a good pratice to copy the struct inside the function,
			// so side effects are prevented if the original struct changes during the time
			type Person struct {
				ID        uint      `gorm:"primary_key"`
				CreatedAt time.Time `gorm:"index"`
				UpdatedAt time.Time
				Name      string
			}
			return tx.AutoMigrate(&Person{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable("people").Error
		},
	})
}
