package postgres

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
)

type execLogger struct {
	logger *log.Logger
	execer sqlx.Execer
}

func (el *execLogger) Exec(query string, args ...interface{}) (sql.Result, error) {
	el.logger.Print(query, args)
	return el.execer.Exec(query, args...)
}
