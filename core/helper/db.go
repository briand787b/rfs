package helper

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

var (
	// ErrNoIDsSingleInsert occurs when no rows are returned from the insertion select
	ErrNoIDsSingleInsert = errors.New("no rows returned from single insertion")
	// ErrTooManyIDsSingleInsert occurs when too many rows are returned from single insertion
	ErrTooManyIDsSingleInsert = errors.New(
		"more than the expected number of rows returned from single insertion")
)

// InsertSingle is a helper function that performs the reception of the returning id
// and the checking of the length of the array that it is returned in.
//
// NOTE: please elaborate on this function later, e.g. what can be passed to it, etc...
func InsertSingle(q sqlx.Queryer, query string, i IDer, args ...interface{}) error {
	var ids []int
	if err := sqlx.Select(q, &ids, query, args...); err != nil {
		return err
	}

	switch {
	case len(ids) == 0:
		return ErrNoIDsSingleInsert
	case len(ids) > 1:
		fmt.Printf("%v ids returned from InsertSingle, expected 1\n", len(ids))
		return ErrTooManyIDsSingleInsert
	case len(ids) == 1:
		return nil
	default:
		return errors.New("unknown error occurred")
	}
}
