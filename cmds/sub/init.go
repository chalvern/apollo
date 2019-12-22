package sub

import (
	"github.com/chalvern/apollo/migrations"
	"github.com/chalvern/sugar"
	"github.com/urfave/cli"
)

var (
	logger *sugar.Logger = sugar.NewLoggerOf("sub_cmd")
)

// Init initial all Sub CMDs
func Init(app *cli.App) {
	app.Commands = []cli.Command{
		commandOfHelloSub(),
		commandOfMigrate(migrations.DefaultMigrator),
	}
}
