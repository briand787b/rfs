package helper

// IDer represents any struct that contains an ID field
type IDer interface {
	GetID() int
	SetID(int)
}
