package postgres

import (
	"context"
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
)

type execContextLogger struct {
	logger        *log.Logger
	execerContext sqlx.ExecerContext
}

func (ecl *execContextLogger) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	ecl.logger.Print(query, args)
	return ecl.execerContext.ExecContext(ctx, query, args)
}
