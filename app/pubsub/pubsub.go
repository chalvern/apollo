package pubsub

import (
	"github.com/chalvern/gochan"
	"github.com/chalvern/sugar"
)

var (
	// Dispatcher 基本的
	defaultDispatcher *gochan.Dispatcher
)

// Init 初始化
func Init() {
	gochan.SetLogger(sugar.NewLoggerOf("pubsub"))

	defaultDispatcher = gochan.NewDispatcher(3, 10)
}

// Dispatch 分发
func Dispatch(objID int, task gochan.TaskFunc) {
	defaultDispatcher.Dispatch(objID, task)
}
