package sqlbuilder

import (
	"GoDBConnection/models/queries/params"
	"fmt"
)

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
	// TO DO: Add numbers to the placeholders in the query
	// Input: "SELECT * FROM users WHERE id = $ and name = $"
	// Output: "SELECT * FROM users WHERE id = $1 and name = $2"
	addIndexPlaceholders(&query)
	return query, args
}

func addIndexPlaceholders(query *string) {
	index := 1
	for i := 0; i < len(*query); i++ {
		if (*query)[i] == '$' {
			*query = (*query)[:i+1] + fmt.Sprint(index) + (*query)[i+1:]
			index++
		}
	}
}
