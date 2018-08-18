package postgres

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"

	// import the postgresql database driver
	_ "github.com/lib/pq"
)

const (
	dbNameEnvVar = "RFS_DATABASE_NAME"
	dbUserEnvVar = "RFS_DATABASE_USER"
	dbPassEnvVar = "RFS_DATABASE_PASS"
)

var (
	db *sqlx.DB
)

func init() {
	// connect to postgresql
	dbName := os.Getenv(dbNameEnvVar)
	if dbName == "" {
		fmt.Println("WARNING: database name is empty")
	}

	dbUser := os.Getenv(dbUserEnvVar)
	if dbUser == "" {
		fmt.Println("WARNING: database user name is empty")
	}

	dbPass := os.Getenv(dbPassEnvVar)
	if dbPass == "" {
		fmt.Println("WARNING: database password is empty")
	}

	var err error
	if db, err = sqlx.Connect("postgres", fmt.Sprintf("sslmode=disable dbname=%s user=%s password=%s",
		dbName,
		dbUser,
		dbPass,
	)); err != nil {
		fmt.Println("WARNING: database connection failed: ", err)
	}
}

// GetDB returns the Postgres database on which to run queries
//
// No need to return any TX's because any helper functions i create
// here will be redundant
func GetDB() *sqlx.DB {
	return db
}
