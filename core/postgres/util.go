package postgres

import "github.com/briand787b/rfs/core/rfslog"

func logQuery(l rfslog.Logger, qry string, args ...interface{}) {
	l.Infow("[PG QUERY]",
		"query", qry,
		"args", args,
	)
}
