package params

type QueryParams interface {
	// AddPlaceholders adds placeholders to the query string. This is currently used for "SELECT" keyword.
	AddPlaceholders(query *string, columns []string) error
	// AddSetColumns adds columns that belongs to the "SET" keyword to the query string.
	AddSetColumns(query *string, columns []string) error
	// AddConditions adds conditions to the query string. This is currently used for "WHERE" keyword.
	AddConditions(query *string, conditions []string) error
}
