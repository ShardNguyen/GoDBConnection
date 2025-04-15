package database

import (
	connections "GoDBConnection/models/connections"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	Connection *connections.PostgresqlConnection
}

func (pg *PostgresDB) InsertRow(query string, args ...interface{}) (int64, error) {
	result, err := pg.Connection.GetDB().Exec(query, args...)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
