package models

// MediaStore is the interface that must be satisfied
// to store metadata about media
type MediaStore interface {
	GetByID(int) (*Media, error)
	Save(*Media) error
}
