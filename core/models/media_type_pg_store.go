package models

import (
	"context"
	"fmt"

	"github.com/briand787b/rfs/core/postgres"
	"github.com/briand787b/rfs/core/rfslog"

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

func (mtps *mediaTypePGStore) GetByID(ctx context.Context, id int) (*MediaType, error) {
	var mtRec MediaType
	if err := sqlx.GetContext(ctx, mtps.db, &mtRec, `
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

func (mtps *mediaTypePGStore) Update(ctx context.Context, mt *MediaType) error {
	qry, args, err := sqlx.Named(`UPDATE media_types SET name = :name WHERE id = :id RETURNING id;`, *mt)
	if err != nil {
		return errors.Wrap(err, "failed to build named query")
	}

	qry = sqlx.Rebind(sqlx.DOLLAR, qry)

	var id int
	if err := sqlx.GetContext(ctx, mtps.db, &id, qry, args...); err != nil {
		return errors.Wrap(err, "failed to execute query")
	}

	return nil
}

func (mtps *mediaTypePGStore) Delete(ctx context.Context, id int) error {
	var delID int
	if err := sqlx.GetContext(ctx, mtps.db, &delID, `
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
