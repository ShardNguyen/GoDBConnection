package queries

type QueryBuilder interface {
	Select(columns ...string) QueryBuilder
	From(table string) QueryBuilder
	Where(condition string, args ...any) QueryBuilder

	InsertInto(table string, columns ...string) QueryBuilder
	Values(values ...any) QueryBuilder

	Update(table string) QueryBuilder
	Set(column string, value any) QueryBuilder

	Build() (string, []any)
}
