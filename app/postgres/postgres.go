package postgres

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"

	// import the postgresql database driver
	_ "github.com/lib/pq"
)

const (
	dbHostEnvVar = "RFS_DATABASE_HOST"
	dbNameEnvVar = "RFS_DATABASE_NAME"
	dbUserEnvVar = "RFS_DATABASE_USER"
	dbPassEnvVar = "RFS_DATABASE_PASS"
)

var (
	db *sqlx.DB
)

func init() {
	// connect to postgresql
	dbHost := os.Getenv(dbHostEnvVar)
	if dbHost == "" {
		fmt.Println("WARNING: databse host is empty")
	}

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

	connStr := fmt.Sprintf("sslmode=disable host=%s dbname=%s user=%s password=%s",
		dbHost,
		dbName,
		dbUser,
		dbPass,
	)

	fmt.Println("db connection string: ", connStr)

	var err error
	if db, err = sqlx.Connect("postgres", connStr); err != nil {
		fmt.Println("WARNING: database connection failed: ", err)
	} else {
		fmt.Println("connected to postgres")
	}
}

// GetDB returns the Postgres database on which to run queries
//
// No need to return any TX's because any helper functions I create
// here will be redundant
func GetDB() *sqlx.DB {
	return db
}
