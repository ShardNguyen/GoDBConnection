package params

import (
	"GoDBConnection/helpers"
	"fmt"
)

type SQLServerParams struct {
}

func NewSQLServerParams() *SQLServerParams {
	return &SQLServerParams{}
}

func (s *SQLServerParams) AddPlaceholders(query *string, columns []string) error {
	if err := helpers.CheckInput(query, columns); err != nil {
		return err
	}

	*query += "@" + fmt.Sprint(columns[0])

	for i := 1; i < len(columns); i++ {
		*query += ", @" + fmt.Sprint(columns[1])
	}

	return nil
}

func (s *SQLServerParams) AddSetColumns(query *string, columns []string) error {
	if err := helpers.CheckInput(query, columns); err != nil {
		return err
	}

	*query += columns[0] + " = @" + fmt.Sprint(columns[0])

	for i := 1; i < len(columns); i++ {
		*query += ", " + columns[i] + " = @" + fmt.Sprint(columns[i])
	}

	return nil
}

func (s *SQLServerParams) AddConditions(query *string, conditions []string) error {
	if err := helpers.CheckInput(query, conditions); err != nil {
		return err
	}

	*query += conditions[0] + " @" + fmt.Sprint(conditions[0])

	for i := 1; i < len(conditions); i++ {
		*query += ", " + conditions[i] + " @" + fmt.Sprint(conditions[i])
	}

	return nil
}
