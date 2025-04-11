package queries

type QueryBuilder interface {
	Select(columns ...string) QueryBuilder
	From(table string) QueryBuilder
	Where(condition string, args ...any) QueryBuilder

	InsertInto(table string, columns ...string) QueryBuilder
	Values(values ...any) QueryBuilder

	BuildSelect() (string, []any)
}
