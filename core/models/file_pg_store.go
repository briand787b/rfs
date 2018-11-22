package models

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type filePGStore struct {
	db *sqlx.DB
}

func (fps *filePGStore) GetByID(int) (*File, error) {
	return nil, errors.New("NOT IMPLEMENTED")
}

func (fps *filePGStore) Save() error {
	return errors.New("NOT IMPLEMENTED")
}

func (fps *filePGStore) Media() (*Media, error) {
	return nil, errors.New("NOT IMPLEMENTED")
}
