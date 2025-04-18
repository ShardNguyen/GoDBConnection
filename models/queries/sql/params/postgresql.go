package params

import "GoDBConnection/helpers"

type PostgresSQLQueryParams struct {
}

func NewPostgresSQLQueryParams() *PostgresSQLQueryParams {
	return &PostgresSQLQueryParams{}
}

func (p *PostgresSQLQueryParams) AddPlaceholders(query *string, columns []string) error {
	if err := helpers.CheckInput(query, columns); err != nil {
		return err
	}

	*query += "$"

	for i := 2; i <= len(columns); i++ {
		*query += ", $"
	}

	return nil
}

func (p *PostgresSQLQueryParams) AddSetColumns(query *string, columns []string) error {
	if err := helpers.CheckInput(query, columns); err != nil {
		return err
	}

	*query += columns[0] + " = $"

	for i := 1; i < len(columns); i++ {
		*query += ", " + columns[i] + " = $"
	}

	return nil
}

func (p *PostgresSQLQueryParams) AddConditions(query *string, conditions []string) error {
	if err := helpers.CheckInput(query, conditions); err != nil {
		return err
	}

	*query += conditions[0] + "$"

	for i := 1; i < len(conditions); i++ {
		*query += ", " + conditions[i] + " = $"
	}

	return nil
}
