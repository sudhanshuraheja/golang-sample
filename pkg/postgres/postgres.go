package postgres

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver
	"github.com/sudhanshuraheja/golang-sample/pkg/logger"
)

const connMaxLifetime = 30 * time.Minute

// NewPostgres - connect to a new postgres server
func NewPostgres(logger logger.Logger, url string, maxOpenConns int) *sqlx.DB {
	db, err := sqlx.Open("postgres", url)
	if err != nil {
		logger.Fatalln("Could not connect to database: %s", err)
	}

	if err = db.Ping(); err != nil {
		logger.Fatalln("Ping to the database failed: %s on connString %s", err, url)
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifetime)
	logger.Debugln("Connected to database")

	return db
}
