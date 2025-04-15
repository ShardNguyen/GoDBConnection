package queries

type QueryBuilder interface {
	Select(columns ...string) QueryBuilder
	From(table string) QueryBuilder
	Where(condition string, args ...any) QueryBuilder

	InsertInto(table string, columns ...string) QueryBuilder
	Values(values ...any) QueryBuilder

	Update(table string) QueryBuilder
	Set(column string, value any) QueryBuilder

	Delete() QueryBuilder

	// This function builds and returns the SQL query and arguments.
	// It will return nothing if the query is not valid.
	Build() (string, []any)
}
