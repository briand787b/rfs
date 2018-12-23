package postgres

import (
	"database/sql"

	"github.com/briand787b/rfs/core/rfslog"
	"github.com/jmoiron/sqlx"
)

type execLogger struct {
	logger rfslog.Logger
	execer sqlx.Execer
}

func (el *execLogger) Exec(query string, args ...interface{}) (sql.Result, error) {
	logQuery(el.logger, query, args)
	return el.execer.Exec(query, args...)
}
