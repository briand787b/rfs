package models

import (
	"fmt"
	"testing"

	"github.com/briand787b/rfs/core/postgres"

	"github.com/jmoiron/sqlx"
)

// TestStruct holds basic information about the model that
// should be created
type TestStruct struct {
	ParentIDs map[int]int // map[<struct enum value>]<struct ID>
	CF        *CleanFunc
	DB        *sqlx.DB
	T         *testing.T
	Count     int
}

// NewTestStruct creates and returns a new TestStruct initialized with
// default values unless passed as a parameter
func NewTestStruct(db *sqlx.DB, t *testing.T, count int) *TestStruct {
	return &TestStruct{
		ParentIDs: make(map[int]int),
		CF:        NewCleanFunc(func() { fmt.Println("Database Cleansed...") }),
		DB:        db,
		T:         t,
		Count:     count,
	}
}

// NewTestStructSimple creates a new TestStruct with default DB and count of 1
// It can do this because it sets up the database/redis as well
func NewTestStructSimple(t *testing.T) *TestStruct {
	return NewTestStruct(postgres.GetDB(), t, 1)
}

// Fatal is a wrapper around testing.T.Fatal that defers
// a call to clean the database before failing the test
func (ts *TestStruct) Fatal(msg string, args ...interface{}) {
	defer (*ts.CF)() // should this be called?  its
	// causing issuesw where failing tests defer cleaning up twice
	ts.T.Fatal(msg, args)
}

// Fatalf is the same as Fatal, but with a formattable message
func (ts *TestStruct) Fatalf(msg string, args ...interface{}) {
	ts.Fatal(fmt.Sprintf(msg, args...))
}
