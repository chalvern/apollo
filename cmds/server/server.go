package server

import (
	"context"

	"github.com/chalvern/sugar"
)

// SimpleThread can execute
type SimpleThread struct {
	Name   string
	Thread func(ctx context.Context)
}

var (
	// Threads would be running in separately goroutines
	Threads []SimpleThread
)

func init() {
	Threads = append(Threads,
		SimpleThread{
			Name: "hello_thread",
			Thread: func(ctx context.Context) {
				sugar.Info("hello Thread")
			}},
	)
}
