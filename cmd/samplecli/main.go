package main

import (
	"os"

	"github.com/urfave/cli"

	"github.com/sudhanshuraheja/golang-sample/pkg/config"
	"github.com/sudhanshuraheja/golang-sample/pkg/logger"
	"github.com/sudhanshuraheja/golang-sample/pkg/postgres"
)

func main() {
	config.Init()
	logger.Init()

	logger.Infoln("Sample CLI")
	Init()
}

// Init : start the cli wrapper
func Init() *cli.App {
	app := cli.NewApp()
	app.Name = config.Name()
	app.Version = config.Version()
	app.Usage = "this service is a sample golang service"

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start the service",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:  "migrate",
			Usage: "run database migrations",
			Action: func(c *cli.Context) error {
				return postgres.RunDatabaseMigrations()
			},
		},
		{
			Name:  "rollback",
			Usage: "rollback the latest database migration",
			Action: func(c *cli.Context) error {
				return postgres.RollbackDatabaseMigration()
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}

	return app
}
