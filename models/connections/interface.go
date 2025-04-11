package connection

import "database/sql"

// Connection is a struct that represents a connection to a database.
type IConnection interface {
	GetDB() *sql.DB
	Close() error
}
