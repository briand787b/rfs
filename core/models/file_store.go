package models

// FileStore is the common interface that any object
// which stores metadata about files (not the files
// themselves) must implement
type FileStore interface {
	GetByID(int) (*File, error)
	Save(*File) error
	GetMedia(*File) (*Media, error)
}
