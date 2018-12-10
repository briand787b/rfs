package postgres

import (
	"context"
	"database/sql"

	"github.com/briand787b/rfs/core/log"

	"github.com/jmoiron/sqlx"
)

type execContextLogger struct {
	logger        log.Logger
	execerContext sqlx.ExecerContext
}

func (ecl *execContextLogger) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	logQuery(ecl.logger, query, args)
	return ecl.execerContext.ExecContext(ctx, query, args)
}
