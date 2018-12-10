package models

import (
	"github.com/briand787b/rfs/core/log"

	"github.com/jmoiron/sqlx"

	"github.com/briand787b/rfs/core/postgres"
)

type testModelDBStore struct {
	l  log.Logger
	DB postgres.ExtFull
}

// NewPostgresTestModelDBStore instantiates a new TestModelStore implemented
// by a postgresql database
func NewPostgresTestModelDBStore(l log.Logger) TestModelStore {
	return &testModelDBStore{
		l:  l,
		DB: postgres.GetExtFull(l),
	}
}

func (s *testModelDBStore) GetAll() ([]TestModel, error) {
	var tms []TestModel
	if err := sqlx.Select(s.DB, &tms, `
		SELECT
			id AS ID,
			name AS Name
		FROM
			test_models;`,
	); err != nil {
		return nil, err
	}

	if tms == nil {
		tms = []TestModel{}
	}

	return tms, nil
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
