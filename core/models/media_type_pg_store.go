package models

import (
	"context"
	"fmt"

	"github.com/briand787b/rfs/core/rfslog"

	"github.com/briand787b/rfs/core/postgres"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type mediaTypePGStore struct {
	l  rfslog.Logger
	db postgres.ExtFull
}

// NewMediaTypePGStore returns a MediaTypeStore backed by Postgresql
func NewMediaTypePGStore(l rfslog.Logger, db postgres.ExtFull) MediaTypeStore {
	return &mediaTypePGStore{l: l, db: db}
}

func (mtps *mediaTypePGStore) GetByID(id int) (*MediaType, error) {
	// delete me
	fmt.Println("DEBUG: in mediaTypePGStore.GetByID")

	var mtRec MediaType
	if err := sqlx.Get(mtps.db, &mtRec, `
		SELECT
			*
		FROM
			media_types
		WHERE
			id = $1;`,
		id,
	); err != nil {
		fmt.Println("error: ", err)
		return nil, errors.Wrap(err, "failed to execute query")
	}

	fmt.Println("NO ERROR")
	return &mtRec, nil
}

func (mtps *mediaTypePGStore) GetAll(ctx context.Context, skip int, take int) ([]MediaType, error) {
	var mts []MediaType
	if err := sqlx.SelectContext(ctx, mtps.db, &mts, `
		SELECT
			*
		FROM
			media_types
		ORDER BY 
			id
		OFFSET 
			$1
		LIMIT
			$2;`,
		skip,
		take,
	); err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}

	fmt.Println("NO ERROR")
	return mts, nil
}

func (mtps *mediaTypePGStore) Insert(ctx context.Context, mt *MediaType) error {
	var saveID int
	if err := sqlx.GetContext(ctx, mtps.db, &saveID, `
		INSERT INTO media_types
		(
			name
		)
		VALUES
		(
			$1
		)
		RETURNING id;`,
		mt.Name,
	); err != nil {
		return errors.Wrap(err, "failed to execute query")
	}

	mt.ID = saveID
	return nil
}

func (mtps *mediaTypePGStore) Update(mt *MediaType) error {
	return errors.New("NOT IMPLEMENTED")
}

func (mtps *mediaTypePGStore) Delete(id int) error {
	var delID int
	if err := sqlx.Get(mtps.db, &delID, `
		DELETE FROM media_types
		WHERE
			id = $1
		RETURNING id;`,
		id,
	); err != nil {
		return errors.Wrap(err, "could not execute query")
	}

	if delID == 0 {
		return errors.New("row was not actually deleted")
	}

	return nil
}
