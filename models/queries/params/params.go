package params

type QueryParams interface {
	AddPlaceholders(query *string, columns []string)
	AddSetColumns(query *string, columns []string)
	AddConditions(query *string, conditions []string)
}
