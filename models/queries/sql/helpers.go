package sqlbuilder

import (
	"strings"
)

func AddTableToSQLQueryString(query *string, table string) {
	*query += table
}

func AddColumnsToSQLQueryString(query *string, columns []string) {
	if len(columns) == 0 {
		*query += "*"
	}

	if len(columns) > 0 {
		*query += strings.Join(columns, ", ")
	}
}
