package models

// MediaTypeStore is the interface that all storers of MediaType must implement
type MediaTypeStore interface {
    GetByID(int) (*MediaType, error)
    Save(*MediaType) error
    Delete(int) error
}

