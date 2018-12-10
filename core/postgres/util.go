package postgres

import (
	"github.com/briand787b/rfs/core/log"
)

func logQuery(l log.Logger, qry string, args ...interface{}) {
	l.Infow("[PG QUERY]",
		"query", qry,
		"args", args,
	)
}
