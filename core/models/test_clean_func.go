package models

// CleanFunc is a function that takes no arguments and
// returns no values.  It is intended to represent
// any function that serves to clean up a database
// after insertion of some model(s)
//
// The most valuable feature of the CleanFunc is that
// additional operations can be added to it through
// a pointer method receiver.  This means that only
// one call to defer is required to defer every
// operation that needs to be cleaned after its
// initial calling
type CleanFunc func()

// NewCleanFunc returns a pointer to a cleaner,
// taking in a bare function
func NewCleanFunc(f func()) *CleanFunc {
	c := CleanFunc(f)
	return &c
}

// Add adds a new operation to the internals of
// the cleaning function method receiver
//
// The provided function is executed BEFORE
// the existing operations.  The order of
// execution from the Cleaners perspective
// can be thought of as a stack
func (c *CleanFunc) Add(f func()) {
	*c = CleanFunc(func() {
		f()
		(*c)()
	})
}
