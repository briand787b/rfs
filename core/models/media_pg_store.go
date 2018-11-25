package models

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type mediaPGStore struct {
	db *sqlx.DB
}

// NewMediaPGStore returns a MediaStore backed by Postgresql
func NewMediaPGStore(db *sqlx.DB) MediaStore { return &mediaPGStore{db: db} }

func (mps *mediaPGStore) GetByID(id int) (m *Media, err error) {
	if err = mps.db.Get(m, `
		SELECT
			id AS ID,
			name AS Name,
			parent_id AS ParentID,
			feature_file_id AS FeatureFileID,
			release_year AS ReleaseYear
		FROM
			media
		WHERE
			id = $1;`,
		id,
	); err != nil {
		err = errors.Wrap(err, "failed to execute query")
	}

	return
}

func (mps *mediaPGStore) Save(m *Media) (err error) {
	var retID int
	if err = mps.db.Get(&retID, `
		INSERT INTO media
		(
			name,
			parent_id,
			feature_file_id,
			release_year
		)
		VALUES
		(
			$1,
			$2,
			$3,
			$4
		)
		RETURNING id;`,
		m.Name,
		m.ParentID,
		m.FeatureFileID,
		m.ReleaseYear,
	); err != nil {
		err = errors.Wrap(err, "failed to execute query")
	}

	return
}

func (mps *mediaPGStore) Delete(id int) (err error) {
	if _, err := mps.db.Exec(`
		DELETE FROM media
		WHERE
			id = $1;`,
		id,
	); err != nil {
		err = errors.Wrap(err, "could not execute query")
	}

	return
}
