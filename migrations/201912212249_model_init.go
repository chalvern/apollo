package migrations

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func init() {
	migrations = append(migrations, &gormigrate.Migration{

		ID: "201912212249",
		Migrate: func(tx *gorm.DB) error {
			// it's a good pratice to copy the struct inside the function,
			// so side effects are prevented if the original struct changes during the time
			type Person struct {
				gorm.Model
				Name string
			}
			return tx.AutoMigrate(&Person{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable("people").Error
		},
	})
}
