package initializer

import (
	"github.com/chalvern/sugar"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitSugarWithPath initializer Sugar with log's path
func InitSugarWithPath(relativePath string) {
	viperInitializedCheck()

	var config zap.Config
	if env := viper.GetString("core.env"); env == "develop" {
		config = zap.NewDevelopmentConfig()
		// config.Encoding = "json"
		config.OutputPaths = []string{"stderr", relativePath + "development.log"}
		config.ErrorOutputPaths = []string{"stderr", relativePath + "development_err.log"}
	} else {
		config = zap.NewProductionConfig()
		config.OutputPaths = []string{relativePath + env + ".log"}
		config.ErrorOutputPaths = []string{relativePath + env + "_err.log"}
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	}
	sugar.SetSugar(&config)
}
