package models

import (
	"fmt"

	"github.com/briand787b/rfs/core/postgres"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type mediaTypePGStore struct {
	db postgres.ExtFull
}

// NewMediaTypePGStore returns a MediaTypeStore backed by Postgresql
func NewMediaTypePGStore(db postgres.ExtFull) MediaTypeStore {
	return &mediaTypePGStore{db: db}
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

func (mtps *mediaTypePGStore) Save(mt *MediaType) error {
	var saveID int
	if err := sqlx.Get(mtps.db, &saveID, `
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
