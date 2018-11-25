package models

import (
	"log"
	"os"
	"testing"

	"github.com/briand787b/rfs/core/postgres"
)

// Table enums
//
// These are snake-cased on purpose to avoid
// any naming collisions with variables where
// they are used
const (
	FILE = iota
	MEDIA
	MEDIA_TYPE
	NETWORK
	SERVER
)

var (
	fileStore      *filePGStore
	mediaStore     *mediaPGStore
	mediaTypeStore *mediaTypePGStore
)

// TestMain controls all common setup and tear down code for model tests
func TestMain(m *testing.M) {
	log.Println("setting up environment for tests...")
	SetUpDependencies()

	log.Println("running test")
	exitStatus := m.Run()

	log.Println("tearing down environment after tests...")
	// put teardown code here

	os.Exit(exitStatus)
}

// SetUpDependencies initializes any required dependencies.  Any
// failure results in the process exiting with status of 1
func SetUpDependencies() {
	db := postgres.GetDB()
	if db == nil {
		log.Println("ERROR: database is nil")
		os.Exit(1)
	}

	fileStore = &filePGStore{db: db}
	mediaStore = &mediaPGStore{db: db}
	mediaTypeStore = NewMediaTypePGStore(db).(*mediaTypePGStore)
}
