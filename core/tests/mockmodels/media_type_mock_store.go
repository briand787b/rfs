package mockmodels

import (
	"context"

	"github.com/briand787b/rfs/core/models"
)

type MediaTypeMockStore struct {
	GetByIDValue *models.MediaType
	GetByIDError error
	GetAllValue  []models.MediaType
	GetAllError  error
	InsertError  error
	UpdateError  error
	DeleteError  error
}

func (mtms *MediaTypeMockStore) GetByID(context.Context, int) (*models.MediaType, error) {
	return mtms.GetByIDValue, mtms.GetByIDError
}
func (mtms *MediaTypeMockStore) GetAll(context.Context, int, int) ([]models.MediaType, error) {
	return mtms.GetAllValue, mtms.GetAllError
}

func (mtms *MediaTypeMockStore) Insert(context.Context, *models.MediaType) error {
	return mtms.InsertError
}
func (mtms *MediaTypeMockStore) Update(context.Context, *models.MediaType) error {
	return mtms.UpdateError
}
func (mtms *MediaTypeMockStore) Delete(context.Context, int) error {
	return mtms.DeleteError
}
