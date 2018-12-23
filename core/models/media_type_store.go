package models

import "context"

// MediaTypeStore is the interface that all storers of MediaType must implement
type MediaTypeStore interface {
	GetByID(int) (*MediaType, error)
	GetAll(context.Context, int, int) ([]MediaType, error)
	Insert(*MediaType) error
	Update(*MediaType) error
	Delete(int) error
}
