package connection

import "fmt"

func GetConnection(dbmsType string) (IConnection, error) {
	switch dbmsType {
	case "postgresql":
		return NewPostgresqlConnection()
	default:
		return nil, fmt.Errorf("unsupported DBMS type: %s", dbmsType)
	}
}
