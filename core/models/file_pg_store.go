package models

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type filePGStore struct {
	db *sqlx.DB
}

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

func (fps *filePGStore) Save() error {
	return errors.New("NOT IMPLEMENTED")
}

func (fps *filePGStore) Media() (*Media, error) {
	return nil, errors.New("NOT IMPLEMENTED")
}
