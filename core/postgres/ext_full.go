package postgres

import (
	"io"
	"log"

	"github.com/jmoiron/sqlx"
)

// ExtFull is the interface that abstracts all the querying
// and executing functions of the sqlx package.  It is satisfied
// by the sqlx.DB type.
type ExtFull interface {
	binder
	sqlx.Execer
	sqlx.ExecerContext
	sqlx.Queryer
	sqlx.QueryerContext
}

// GetExtFull returns an implementation of ExtFull that uses postgres
func GetExtFull(logOut io.Writer) ExtFull {
	if logOut == nil {
		return db
	}

	l := log.New(logOut, "[QUERY] ", log.LstdFlags)
	return struct {
		binder
		sqlx.Execer
		sqlx.ExecerContext
		sqlx.Queryer
		sqlx.QueryerContext
	}{
		db,
		&execLogger{
			logger: l,
			execer: db,
		},
		&execContextLogger{
			logger:        l,
			execerContext: db,
		},
		&queryLogger{
			logger:  l,
			queryer: db,
		},
		&queryContextLogger{
			logger:         l,
			queryerContext: db,
		},
	}
}
