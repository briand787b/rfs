package helper

// IDer represents any struct that contains a
// retrievalbe ID field
type IDer interface {
	GetID() int
	SetID(int)
}
