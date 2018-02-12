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
	config := config.NewConfig([]string{".", "../.."})
	logger := logger.NewLogger(config, os.Stdout)
	ctx := appcontext.NewAppContext(config, logger)
	db := postgres.NewPostgres(logger, config.Database().ConnectionURL(), config.Database().MaxPoolSize())
	server := server.NewServer(ctx, db)

	logger.Infoln("Starting sample-cli")

	app := cli.NewApp()
	app.Name = "sample"
	app.Version = "0.0.1"
	app.Usage = "this service is a sample golang service"

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start the service",
			Action: func(c *cli.Context) error {
				return server.Start()
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
}
