package initializer

import (
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
)

var (
	viperInitialized = false
)

// InitViperWithFile for viper
// more see https://github.com/spf13/viper
// 初始化配置工具 viper
func InitViperWithFile(configFile string) {
	jww.SetStdoutThreshold(jww.LevelInfo)

	viper.SetConfigType("yaml")

	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// default value
	viper.SetDefault("core.env", "production")
	viper.SetDefault("core.monitor_addr", ":9999")

	viperInitialized = true
}

func viperInitializedCheck() {
	if !viperInitialized {
		panic("viper must be initialized first!")
	}
}
