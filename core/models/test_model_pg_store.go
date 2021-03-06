package models

import (
	"github.com/jmoiron/sqlx"

	"github.com/briand787b/rfs/core/postgres"
	"github.com/briand787b/rfs/core/rfslog"
)

type testModelDBStore struct {
	l  rfslog.Logger
	DB postgres.ExtFull
}

// NewPostgresTestModelDBStore instantiates a new TestModelStore implemented
// by a postgresql database
func NewPostgresTestModelDBStore(l rfslog.Logger) TestModelStore {
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
