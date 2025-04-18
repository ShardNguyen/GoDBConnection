package sqldbsv

// Database is an interface that defines methods for interacting with a database.
type Database interface {
	CreateTable() error
	DropTable() error
	InsertRow() error
	UpdateRow() error
	DeleteRow() error
	SelectRow() error
	SelectAllRows() error
}
