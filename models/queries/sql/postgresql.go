package sqlqueries

import (
	"GoDBConnection/models/queries/sql/params"
	"fmt"
)

type PostgresSQLQueryBuilder struct {
	queryBuilder DefaultSQLQueryBuilder
}

func NewPostgresSQLQueryBuilder() *PostgresSQLQueryBuilder {
	NewPostgresSQLQueryBuilder := &PostgresSQLQueryBuilder{
		queryBuilder: *NewSQLQueryBuilder(),
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

// Postgres's Build function constructs the SQL query and its arguments then adds index to the placeholders.
func (qb *PostgresSQLQueryBuilder) Build() (*string, []any, error) {
	query, args, err := qb.queryBuilder.Build()

	if err != nil {
		return nil, nil, err
	}

	addIndexToPlaceholders(query)
	return query, args, nil
}

func addIndexToPlaceholders(query *string) {
	index := 1
	// DO NOT CHANGE THIS TO IN RANGE
	for i := 0; i < len(*query); i++ {
		if (*query)[i] == '$' {
			*query = (*query)[:i+1] + fmt.Sprint(index) + (*query)[i+1:]
			index++
		}
	}
}
