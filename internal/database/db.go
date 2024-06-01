// db.go
package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/database/postgres"
	_ "github.com/lib/pq"
)

// New initializes and returns a new instance of the database store.
func New(db *postgres.Queries) *Store {
	return &Store{
		DB: db,
	}
}

// Store represents the database store.
type Store struct {
	DB *postgres.Queries
}

// OpenDB opens a connection to the PostgreSQL database.
func OpenDB() (*postgres.Queries, error) {
	dbConn := os.Getenv("DB_URL")

	conn, err := sql.Open("postgres", dbConn)
	if err != nil {
		log.Fatal(err)
	}
	db := postgres.New(conn)
	return db, nil
}
