package main

import (
	"os"

	"github.com/urfave/cli"

	"github.com/sudhanshuraheja/golang-sample/pkg/appcontext"
	"github.com/sudhanshuraheja/golang-sample/pkg/config"
	"github.com/sudhanshuraheja/golang-sample/pkg/logger"
	"github.com/sudhanshuraheja/golang-sample/pkg/postgres"
	"github.com/sudhanshuraheja/golang-sample/pkg/server"
)

func main() {
	config := config.NewConfig()
	logger := logger.NewLogger(config)
	ctx := appcontext.NewAppContext(config, logger)

	logger.Infoln("Starting sample-cli")

	app := cli.NewApp()
	app.Name = config.Name()
	app.Version = config.Version()
	app.Usage = "this service is a sample golang service"

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start the service",
			Action: func(c *cli.Context) error {
				return server.StartAPIServer(ctx)
			},
		},
		{
			Name:  "migrate",
			Usage: "run database migrations",
			Action: func(c *cli.Context) error {
				return postgres.RunDatabaseMigrations(ctx)
			},
		},
		{
			Name:  "rollback",
			Usage: "rollback the latest database migration",
			Action: func(c *cli.Context) error {
				return postgres.RollbackDatabaseMigration(ctx)
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}

	return app
}
