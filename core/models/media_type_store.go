package models

import "context"

// MediaTypeStore is the interface that all storers of MediaType must implement
type MediaTypeStore interface {
	GetByID(context.Context, int) (*MediaType, error)
	GetAll(context.Context, int, int) ([]MediaType, error)
	// delete these once the patterns have been applied elsewhere - media types are static
	Insert(context.Context, *MediaType) error
	Update(context.Context, *MediaType) error
	Delete(context.Context, int) error
}
