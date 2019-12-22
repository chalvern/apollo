package server

import (
	"context"

	"github.com/chalvern/apollo/app"
)

func init() {
	Threads = append(Threads,
		SimpleThread{
			Name: "web_app",
			Thread: func(ctx context.Context) {
				app.Run(ctx)
			}},
	)
}
