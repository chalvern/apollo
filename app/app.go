package app

import (
	"context"
	"github.com/spf13/viper"
)

// Run app run
func Run(ctx context.Context) {
	var (
		addr = viper.GetString("webserver.addr")
		env  = viper.GetString("core.env")
	)

	// 1. production mode for default
	// 1. 默认为生产模式
	if env != "develop" {
		gin.SetMode(gin.ReleaseMode)
	}

}
