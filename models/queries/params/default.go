package params

type DefaultQueryParams struct {
}

func NewDefaultQueryParams() *DefaultQueryParams {
	return &DefaultQueryParams{}
}

func (p *DefaultQueryParams) AddPlaceholders(query *string, columns []string) {
	*query += "?"

	for i := 1; i < len(columns); i++ {
		*query += ", ?"
	}
}

func (p *DefaultQueryParams) AddSetColumns(query *string, columns []string) {
	if len(columns) == 0 {
		return
	}

	*query += columns[0] + " = ?"

	for i := 1; i < len(columns); i++ {
		*query += ", " + columns[i] + " = ?"
	}
}

func (p *DefaultQueryParams) AddConditions(query *string, conditions []string) {
	if len(conditions) == 0 {
		return
	}

	*query += conditions[0] + "?"

	for _, column := range conditions[1:] {
		*query += ", " + column + "?"
	}
}
