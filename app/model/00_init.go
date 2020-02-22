package model

import (
	"time"

	"github.com/chalvern/apollo/configs/initializer"
	"github.com/chalvern/sugar"
	"github.com/jinzhu/gorm"
)

// Model 通用的 Model
type Model struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `gorm:"" json:"-"`
	UpdatedAt time.Time `gorm:"" json:"-"`
}

var (
	mydb   *gorm.DB
	logger *sugar.Logger
)

// Init initial database
func Init() {
	mydb = initializer.DB
	logger = sugar.NewLoggerOf("db")
}

// SetMyDB 设置 db
func SetMyDB(db *gorm.DB) {
	mydb = db
}

// dbArgs 处理变量函数
func dbArgs(db *gorm.DB, args ...interface{}) *gorm.DB {
	if len(args) >= 2 {
		db = db.Where(args[0], args[1:]...)
	} else if len(args) >= 1 {
		db = db.Where(args[0])
	}
	return db
}
