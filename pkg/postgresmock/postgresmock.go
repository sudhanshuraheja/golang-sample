package postgresmock

import (
	"log"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

// NewMockSqlxDB - create a new mock
func NewMockSqlxDB() (*sqlx.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Should not get an error in creating a mock database")
	}
	sqlxDB := sqlx.NewDb(db, "postgres")
	sqlxDB.SetMaxOpenConns(10)
	return sqlxDB, mock
}

// CloseMockSqlxDB - close the mock
func CloseMockSqlxDB(db *sqlx.DB) {
	db.Close()
}
