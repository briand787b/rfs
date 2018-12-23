package postgres

import (
	"context"
	"database/sql"

	"github.com/briand787b/rfs/core/rfslog"
	"github.com/jmoiron/sqlx"
)

type execContextLogger struct {
	logger        rfslog.Logger
	execerContext sqlx.ExecerContext
}

func (ecl *execContextLogger) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	logQuery(ecl.logger, query, args)
	return ecl.execerContext.ExecContext(ctx, query, args)
}
