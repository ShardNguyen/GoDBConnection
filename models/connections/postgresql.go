package connection

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type PostgresqlConnection struct {
	db *sql.DB
}

func NewPostgresqlConnection() (*PostgresqlConnection, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %v", err)
	}

	// Verify the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	log.Println("PostgreSQL connection established successfully")
	return &PostgresqlConnection{db: db}, nil
}

func (pc *PostgresqlConnection) GetDB() *sql.DB {
	return pc.db
}

func (pc *PostgresqlConnection) Close() error {
	if pc.db != nil {
		return pc.db.Close()
	}
	return nil
}
