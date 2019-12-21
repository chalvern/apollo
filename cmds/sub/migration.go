package sub

import (
	"context"

	"github.com/chalvern/apollo/configs/initializer"
	"github.com/chalvern/apollo/migrations"
	"github.com/chalvern/apollo/migrations/template"
	"github.com/chalvern/sugar"
	"github.com/urfave/cli"
)

var (
	migrator migrations.Migrator
)

// CommandOfMigrate command of migrate
func CommandOfMigrate(migratorArg migrations.Migrator) cli.Command {
	migrator = migratorArg
	return cli.Command{
		Name:    "migrate",
		Aliases: []string{"m"},
		Usage:   "migrate all migrations defined",
		Action:  migrate,
		Subcommands: []cli.Command{
			{
				Name:    "createMigration",
				Aliases: []string{"cm"},
				Usage:   "create migration files with 'timestamp_' + '[some_hint_name_you_defined]'",
				Action:  createMigration,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "path, p",
						Value: "./migrations",
						Usage: "the path of migration file",
					},
				},
			},
			{
				Name:    "migrateTo",
				Aliases: []string{"mt"},
				Usage:   "executes all migrations that did not run yet up to the migration that matches `migrationID`",
				Action:  migrateTo,
			},
			{
				Name:    "rollbackLast",
				Aliases: []string{"rl"},
				Usage:   "undo the last migration",
				Action:  rollbackLast,
			},
			{
				Name:    "rollbackTo",
				Aliases: []string{"rt"},
				Usage:   "undoes migrations up to the given migration that matches the `migrationID`.",
				Action:  rollbackTo,
			},
		},
	}
}

func createMigration(c *cli.Context) error {
	args := c.Args()
	if len(args) != 1 {
		sugar.Fatalf("must exactly one arg (hint_name of migration file) defined, args are now %v.", args)
	}
	hintName := args[0]
	return template.CreateNewTable(c.String("path"), hintName)
}

func migrate(c *cli.Context) error {
	sugar.Infof("starting migration with config: %s", c.String("config"))

	// initializer
	initializer.InitMysql(context.Background())

	migrator.Migrate()

	sugar.Info("migration done!")
	return nil
}

func migrateTo(c *cli.Context) error {
	sugar.Info("starting migrateTo with config: %s", c.String("config"))

	// initializer
	initializer.InitMysql(context.Background())

	args := c.Args()
	if len(args) != 1 {
		sugar.Fatalf("one migrationID（timestamp）must be provided")
		return nil
	}

	migrator.MigrateTo(args[0])

	sugar.Infof("migration to %s done!", args[0])
	return nil
}

func rollbackLast(c *cli.Context) error {
	sugar.Info("starting rollbackLast with config: %s", c.String("config"))

	// initializer
	initializer.InitMysql(context.Background())

	migrator.RollbackLast()

	sugar.Info("RollbackLast done!")
	return nil
}

func rollbackTo(c *cli.Context) error {
	sugar.Info("starting rollbackTo with config: %s", c.String("config"))

	// initializer
	initializer.InitMysql(context.Background())

	args := c.Args()
	if len(args) != 1 {
		sugar.Fatalf("one migrationID（timestamp）must be provided")
		return nil
	}

	migrator.RollbackTo(args[0])

	sugar.Infof("RollbackTo to %s done!", args[0])
	return nil
}
