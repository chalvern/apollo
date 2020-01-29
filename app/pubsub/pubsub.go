package pubsub

import (
	"github.com/chalvern/gochan"
	"github.com/chalvern/sugar"
	"github.com/spf13/viper"
)

var (
	// Dispatcher 基本的
	defaultDispatcher *gochan.Dispatcher

	pubsubGochanNum int
	pubsubBufferNum int
)

// Init 初始化
func Init() {
	gochan.SetLogger(sugar.NewLoggerOf("pubsub"))
	pubsubGochanNum = viper.GetInt("pubsub.gochan_num")
	pubsubBufferNum = viper.GetInt("pubsub.buffer_num")

	if pubsubGochanNum == 0 || pubsubBufferNum == 0 {
		pubsubGochanNum = 1
		pubsubBufferNum = 3
	}
	defaultDispatcher = gochan.NewDispatcher(pubsubGochanNum, pubsubBufferNum)
}

// Dispatch 分发
func Dispatch(objID int, task gochan.TaskFunc) {
	defaultDispatcher.Dispatch(objID, task)
}
