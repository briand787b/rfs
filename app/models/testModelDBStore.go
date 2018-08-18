package models

import (
	"os"

	"github.com/jmoiron/sqlx"

	"github.com/briand787b/rfs/app/postgres"
)

type testModelDBStore struct {
	DB postgres.ExtFull
}

// NewPostgresTestModelDBStore instantiates a new TestModelStore implemented
// by a postgresql database
func NewPostgresTestModelDBStore() TestModelStore {
	return &testModelDBStore{
		DB: postgres.GetExtFull(os.Stdout),
	}
}

func (s *testModelDBStore) Save(tm *TestModel) error {
	var ids []int
	if err := sqlx.Select(s.DB, &ids, `
		INSERT INTO test_models 
		(
			name
		)
		VALUES
		(
			$1
		)
		RETURNING id;`,
		tm.Name,
	); err != nil {
		return err
	}

	tm.ID = ids[0]

	// if err := r.Scan(&tm.ID); err != nil {
	// 	return errors.Wrap(err,
	// 		"could not scan returned inserted id",
	// 	)
	// }

	return nil
}
