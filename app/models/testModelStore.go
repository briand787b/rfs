package models

type TestModelStore interface {
	Save(*TestModel) error
	GetAll() ([]TestModel, error)
}
