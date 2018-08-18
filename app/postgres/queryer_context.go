package postgres

import (
	"context"
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
)

type queryContextLogger struct {
	logger         *log.Logger
	queryerContext sqlx.QueryerContext
}

func (qcl *queryContextLogger) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	qcl.logger.Print(query, args)
	return qcl.queryerContext.QueryContext(ctx, query, args...)
}

func (qcl *queryContextLogger) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	qcl.logger.Print(query, args)
	return qcl.queryerContext.QueryxContext(ctx, query, args...)
}

func (qcl *queryContextLogger) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	qcl.logger.Print(query, args)
	return qcl.queryerContext.QueryRowxContext(ctx, query, args...)
}
