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

func (mps *mediaPGStore) GetByID(int) (*Media, error) {
	return nil, errors.New("NOT IMPLEMENTED")
}

func (mps *mediaPGStore) Save(*Media) error {
	return errors.New("NOT IMPLEMENTED")
}
