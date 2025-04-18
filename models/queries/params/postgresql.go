package params

type PostgresSQLQueryParams struct {
}

func NewPostgresSQLQueryParams() *PostgresSQLQueryParams {
	return &PostgresSQLQueryParams{}
}

func (p *PostgresSQLQueryParams) AddPlaceholders(query *string, columns []string) {
	*query += "$"

	for i := 2; i <= len(columns); i++ {
		*query += ", $"
	}
}

func (p *PostgresSQLQueryParams) AddSetColumns(query *string, columns []string) {
	if len(columns) == 0 {
		return
	}

	*query += columns[0] + " = $"

	for i := 1; i < len(columns); i++ {
		*query += ", " + columns[i] + " = $"
	}
}

func (p *PostgresSQLQueryParams) AddConditions(query *string, conditions []string) {
	if len(conditions) == 0 {
		return
	}

	*query += conditions[0] + "$"

	for i := 1; i < len(conditions); i++ {
		*query += ", " + conditions[i] + " = $"
	}
}
