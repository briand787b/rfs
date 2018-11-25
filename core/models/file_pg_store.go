package models

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type filePGStore struct {
	db *sqlx.DB
}

// NewFilePGStore returns a FileStore backed by Postgresql
func NewFilePGStore(db *sqlx.DB) FileStore { return &filePGStore{db: db} }

func (fps *filePGStore) GetByID(id int) (f *File, err error) {
	if err = fps.db.Get(f, `
		SELECT
			id,
			media_id,
			md5_checksum
		FROM
			files
		WHERE
			id = $1;`,
		id,
	); err != nil {
		err = errors.Wrap(err, "failed to execute query")
	}

	return
}

func (fps *filePGStore) Save(f *File) error {
	return errors.New("NOT IMPLEMENTED")
}

func (fps *filePGStore) Delete(id int) error {
	return errors.New("NOT IMPLEMENTED")
}
