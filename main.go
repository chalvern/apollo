package main

import (
	"os"

	"github.com/chalvern/apollo/cmds"
)

func main() {

	app := cmds.AppInit()
	app.Run(os.Args)
}
