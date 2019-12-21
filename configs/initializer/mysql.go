package initializer

import (
	"context"
	"fmt"

	"github.com/chalvern/sugar"

	"github.com/jinzhu/gorm"
	// 用于 gorm 底层使用
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

// DB 数据库实例
var DB *gorm.DB

// InitMysql 初始化数据库
func InitMysql(ctx context.Context) {
	viperInitializedCheck()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.dbname"),
		viper.GetString("database.args"))

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("init database error: " + err.Error())
	}

	if viper.GetBool("database.row_format_dynamic") {
		db = db.Set("gorm:table_options", "ROW_FORMAT=DYNAMIC")
	}
	DB = db

	DB.DB().SetMaxIdleConns(viper.GetInt("database.max_idle_conns"))
	DB.DB().SetMaxOpenConns(viper.GetInt("database.max_open_conns"))
	if viper.GetBool("database.log_mod") {
		DB.LogMode(true)
	}
	go closeMysql(ctx)
}

// InitTestMysql 初始化测试数据库
func InitTestMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		"root", "123456", "127.0.0.1", "3306", "jzb", "charset=utf8mb4&parseTime=True&loc=Local",
	)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("init database error: " + err.Error())
	}
	DB = db
	DB.LogMode(true)
}

func closeMysql(ctx context.Context) {
	<-ctx.Done()
	DB.Close()
	sugar.Info("database closed successfully")
}
