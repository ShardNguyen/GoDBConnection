package sqlqueries

type SQLQueryBuilder interface {
	Select(columns ...string) SQLQueryBuilder
	From(table string) SQLQueryBuilder
	Where(condition string, args ...any) SQLQueryBuilder

	InsertInto(table string, columns ...string) SQLQueryBuilder
	Values(values ...any) SQLQueryBuilder

	Update(table string) SQLQueryBuilder
	Set(column string, value any) SQLQueryBuilder

	Delete() SQLQueryBuilder

	Create() SQLQueryBuilder
	Drop() SQLQueryBuilder
	Table(table string) SQLQueryBuilder
	// This function builds and returns the SQL query and arguments.
	// It will return nothing if the query is not valid.
	Build() (string, []any)
}
