package postgres

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"

	// import the postgresql database driver
	_ "github.com/lib/pq"
)

const (
	dbHostEnvVar = "RFS_DATABASE_HOST"
	dbNameEnvVar = "RFS_DATABASE_NAME"
	dbUserEnvVar = "RFS_DATABASE_USER"
	dbPassEnvVar = "RFS_DATABASE_PASS"
	dbPortEnvVar = "RFS_DATABASE_PORT"
)

var (
	db *sqlx.DB

	connectOnce = &sync.Once{}
)

func connect() {
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

	dbPort := os.Getenv(dbPortEnvVar)
	if dbPort == "" {
		fmt.Println("WARNING: database port is empty")
	}

	connStr := fmt.Sprintf("sslmode=disable host=%s dbname=%s user=%s password=%s port=%s",
		dbHost,
		dbName,
		dbUser,
		dbPass,
		dbPort,
	)

	// DEBUG
	fmt.Println("DEBUG: db connection string: ", connStr)

	var err error
	if db, err = sqlx.Connect("postgres", connStr); err == nil {
		fmt.Println("connected to postgres!")
		return
	}

	go func() {
		for {
			if db, err = sqlx.Connect("postgres", connStr); err == nil {
				fmt.Println("connected to postgres!")
				return
			}

			fmt.Println("WARNING: database connection failed: ", err)
			time.Sleep(50 * time.Millisecond)
			fmt.Println("retrying...")
		}
	}()

}

// GetDB returns the Postgres database on which to run queries
//
// No need to return any TX's because any helper functions I create
// here will be redundant
func GetDB() *sqlx.DB {
	connectOnce.Do(connect)
	return db
}
