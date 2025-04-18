package params

import "fmt"

type SQLServerParams struct {
}

func NewSQLServerParams() *SQLServerParams {
	return &SQLServerParams{}
}

func (s *SQLServerParams) AddPlaceholders(query *string, columns []string) {
	if len(columns) == 0 {
		return
	}

	*query += "@" + fmt.Sprint(columns[0])

	for i := 1; i < len(columns); i++ {
		*query += ", @" + fmt.Sprint(columns[1])
	}
}

func (s *SQLServerParams) AddSetColumns(query *string, columns []string) {
	if len(columns) == 0 {
		return
	}

	*query += columns[0] + " = @" + fmt.Sprint(columns[0])

	for i := 1; i < len(columns); i++ {
		*query += ", " + columns[i] + " = @" + fmt.Sprint(columns[i])
	}
}

func (s *SQLServerParams) AddConditions(query *string, conditions []string) {
	if len(conditions) == 0 {
		return
	}

	*query += conditions[0] + " @" + fmt.Sprint(conditions[0])

	for i := 1; i < len(conditions); i++ {
		*query += ", " + conditions[i] + " @" + fmt.Sprint(conditions[i])
	}
}
