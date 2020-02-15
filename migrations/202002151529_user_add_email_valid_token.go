// Package migrations auto generated file
package migrations

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func init() {
	migrations = append(migrations, &gormigrate.Migration{

		ID: "202002151529",
		Migrate: func(tx *gorm.DB) error {
			// it's a good pratice to copy the struct inside the function,
			// so side effects are prevented if the original struct changes during the time
			type User struct {
				EmailValidToken string `gorm:"type:varchar(50)" json:"-"` // 邮件认证的token
			}
			return tx.AutoMigrate(&User{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Table("users").DropColumn("email_valid_token").Error
		},
	})
}
