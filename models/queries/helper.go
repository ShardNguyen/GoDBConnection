package queries

import "strings"

func AddColumnsToSQLQueryString(query *string, columns []string) {
	if len(columns) > 0 {
		*query += strings.Join(columns, ", ")
	}
}

func AddConditionsToSQLQueryString(query *string, conditions []string) {
	if len(conditions) > 0 {
		*query += " WHERE " + strings.Join(conditions, " AND ")
	}
}
