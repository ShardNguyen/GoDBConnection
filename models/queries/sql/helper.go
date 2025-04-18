package sqlqueries

import (
	"GoDBConnection/helpers"
	"errors"
	"strings"
)

func AddTable(query *string, table string) error {
	if helpers.CheckQueryIsEmpty(query) {
		return errors.New("query is empty")
	}

	if table == "" {
		return errors.New("table is empty")
	}

	*query += table
	return nil
}

func AddColumns(query *string, columns []string) {
	if len(columns) == 0 {
		*query += "*"
	}

	if len(columns) > 0 {
		*query += strings.Join(columns, ", ")
	}
}
