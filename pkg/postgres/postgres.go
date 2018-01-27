package postgres

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver
	"github.com/sudhanshuraheja/golang-sample/pkg/config"
	"github.com/sudhanshuraheja/golang-sample/pkg/logger"
)

const (
	connMaxLifetime = 30 * time.Minute
	defaultTimeout  = 1 * time.Second
)

var database *sqlx.DB

func Init() {
	var err error

	database, err = sqlx.Open("postgres", config.Database().ConnectionString())
	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	} else {
		logger.Debug("Connected to database")
	}

	if err = database.Ping(); err != nil {
		log.Fatalf("Ping to the database failed: %s on connString %s", err, config.Database().ConnectionString())
	}

	database.SetMaxIdleConns(config.Database().MaxPoolSize())
	database.SetMaxOpenConns(config.Database().MaxPoolSize())
	database.SetConnMaxLifetime(connMaxLifetime)
}

func Close() error {
	logger.Debug("Closing the DB connection")
	return database.Close()
}

func Get() *sqlx.DB {
	return database
}
