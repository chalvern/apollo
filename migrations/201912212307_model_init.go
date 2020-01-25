// Package migrations auto generated file
package migrations

import (
	"time"

	"github.com/chalvern/apollo/app/model"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func init() {
	migrations = append(migrations, &gormigrate.Migration{

		ID: "201912212307",
		Migrate: func(tx *gorm.DB) error {
			// it's a good pratice to copy the struct inside the function,
			// so side effects are prevented if the original struct changes during the time
			type User struct {
				model.Model
				// Username 用户名，限定字符数目
				Email    string `gorm:"type:varchar(100);unique_index" json:"email"`
				Password string `gorm:"type:varchar(100);" json:"-"`

				EmailVarified bool `gorm:"" json:"-"` // 邮件已认证

				ResetPasswordKey string     `gorm:"type:varchar(128)" json:"-"` // 重置密码所需要的key
				BannedTime       *time.Time `gorm:""`                           // 被拉黑到什么时候
			}
			return tx.AutoMigrate(&User{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable("users").Error
		},
	})
}
