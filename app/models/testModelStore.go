package models

type TestModelStore interface {
	Save(*TestModel) error
}
