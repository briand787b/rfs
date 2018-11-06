package postgres

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
)

type queryLogger struct {
	logger  *log.Logger
	queryer sqlx.Queryer
}

func (ql *queryLogger) Query(query string, args ...interface{}) (*sql.Rows, error) {
	ql.logger.Print(query, args)
	return ql.queryer.Query(query, args...)
}

func (ql *queryLogger) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	ql.logger.Print(query, args)
	return ql.queryer.Queryx(query, args...)
}

func (ql *queryLogger) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	ql.logger.Print(query, args)
	return ql.queryer.QueryRowx(query, args...)
}
