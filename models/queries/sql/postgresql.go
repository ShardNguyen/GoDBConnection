package sqlbuilder

import "GoDBConnection/models/queries/params"

type PostgresSQLQueryBuilder struct {
	queryBuilder DefaultSQLQueryBuilder
}

func NewPostgresSQLQueryBuilder(table string) *PostgresSQLQueryBuilder {
	NewPostgresSQLQueryBuilder := &PostgresSQLQueryBuilder{
		queryBuilder: *NewSQLQueryBuilder(table),
	}
	NewPostgresSQLQueryBuilder.queryBuilder.param = params.NewPostgresSQLQueryParams()
	return NewPostgresSQLQueryBuilder
}

func (qb *PostgresSQLQueryBuilder) Select(columns ...string) *PostgresSQLQueryBuilder {
	qb.queryBuilder.Select(columns...)
	return qb
}

func (qb *PostgresSQLQueryBuilder) From(table string) *PostgresSQLQueryBuilder {
	qb.queryBuilder.From(table)
	return qb
}

func (qb *PostgresSQLQueryBuilder) Where(condition string, arg any) *PostgresSQLQueryBuilder {
	qb.queryBuilder.Where(condition, arg)
	return qb
}

func (qb *PostgresSQLQueryBuilder) InsertInto(table string, columns ...string) *PostgresSQLQueryBuilder {
	qb.queryBuilder.InsertInto(table, columns...)
	return qb
}

func (qb *PostgresSQLQueryBuilder) Values(values ...any) *PostgresSQLQueryBuilder {
	qb.queryBuilder.Values(values...)
	return qb
}

func (qb *PostgresSQLQueryBuilder) Update(table string) *PostgresSQLQueryBuilder {
	qb.queryBuilder.Update(table)
	return qb
}

func (qb *PostgresSQLQueryBuilder) Set(column string, value any) *PostgresSQLQueryBuilder {
	qb.queryBuilder.Set(column, value)
	return qb
}

func (qb *PostgresSQLQueryBuilder) Delete() *PostgresSQLQueryBuilder {
	qb.queryBuilder.Delete()
	return qb
}

func (qb *PostgresSQLQueryBuilder) Build() (string, []any) {
	query, args := qb.queryBuilder.Build()

	return query, args
}
