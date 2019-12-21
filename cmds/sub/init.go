package sub

import (
	"github.com/urfave/cli"
)

// Init initial all Sub CMDs
func Init(app *cli.App) {
	app.Commands = []cli.Command{
		commandOfHelloSub(),
	}
}
