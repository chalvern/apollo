package sub

import (
	"github.com/chalvern/sugar"

	"github.com/urfave/cli"
)

// commandOfHelloSub return command of helloSub
func commandOfHelloSub() cli.Command {
	return cli.Command{
		Name:    "helloSub",
		Aliases: []string{"hs"},
		Usage:   "sub command example",
		Action:  helloSub,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "flagExample, f",
				Value: "I am flagExample",
				Usage: "for flag example of sub command",
			},
		},
	}
}

func helloSub(c *cli.Context) error {
	args := c.Args()
	if len(args) > 0 {
		sugar.Infof("Args are: %v", args[0])
	}
	sugar.Infof("flagExample Value: %s", c.String("flagExample"))
	sugar.Info("Hello sub command!")
	return nil
}
