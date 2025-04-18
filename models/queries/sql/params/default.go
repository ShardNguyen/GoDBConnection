package params

import "GoDBConnection/helpers"

type DefaultQueryParams struct {
}

func NewDefaultQueryParams() *DefaultQueryParams {
	return &DefaultQueryParams{}
}

func (p *DefaultQueryParams) AddPlaceholders(query *string, columns []string) error {
	if err := helpers.CheckInput(query, columns); err != nil {
		return err
	}

	*query += "?"

	for i := 1; i < len(columns); i++ {
		*query += ", ?"
	}

	return nil
}

func (p *DefaultQueryParams) AddSetColumns(query *string, columns []string) error {
	if err := helpers.CheckInput(query, columns); err != nil {
		return err
	}

	*query += columns[0] + " = ?"

	for i := 1; i < len(columns); i++ {
		*query += ", " + columns[i] + " = ?"
	}

	return nil
}

func (p *DefaultQueryParams) AddConditions(query *string, conditions []string) error {
	if err := helpers.CheckInput(query, conditions); err != nil {
		return err
	}

	*query += conditions[0] + "?"

	for _, column := range conditions[1:] {
		*query += ", " + column + "?"
	}

	return nil
}
