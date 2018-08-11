package app

// MediaStore is the interface that any struct which persists Media must implement
type MediaStore interface {
	AddMediaHTTP(f *Media) error
	MarshalJSON() ([]byte, error)
}
