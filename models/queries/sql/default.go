package sqlbuilder

import (
	"log"

	"GoDBConnection/models/queries/params"
)

const (
	SelectKeyWord = "SELECT "
	InsertKeyWord = "INSERT INTO "
	UpdateKeyWord = "UPDATE "
	DeleteKeyWord = "DELETE "
)

type DefaultSQLQueryBuilder struct {
	table      string
	columns    []string
	conditions []string
	args       []any
	values     []any
	query      string
	param      params.QueryParams
}

func NewSQLQueryBuilder(table string) *DefaultSQLQueryBuilder {
	return &DefaultSQLQueryBuilder{
		table:      table,
		columns:    []string{},
		conditions: []string{},
		args:       []any{},
		values:     []any{},
		query:      "",
		param:      params.NewDefaultQueryParams(),
	}
}

func (qb *DefaultSQLQueryBuilder) Select(columns ...string) *DefaultSQLQueryBuilder {
	qb.query += SelectKeyWord
	qb.columns = columns
	return qb
}

func (qb *DefaultSQLQueryBuilder) From(table string) *DefaultSQLQueryBuilder {
	qb.table = table
	return qb
}

func (qb *DefaultSQLQueryBuilder) Where(condition string, arg any) *DefaultSQLQueryBuilder {
	qb.conditions = append(qb.conditions, condition)
	qb.args = append(qb.args, arg)
	return qb
}

func (qb *DefaultSQLQueryBuilder) InsertInto(table string, columns ...string) *DefaultSQLQueryBuilder {
	qb.query += InsertKeyWord
	qb.table = table
	qb.columns = columns
	return qb
}

func (qb *DefaultSQLQueryBuilder) Values(values ...any) *DefaultSQLQueryBuilder {
	qb.values = values
	return qb
}

func (qb *DefaultSQLQueryBuilder) Update(table string) *DefaultSQLQueryBuilder {
	qb.query += UpdateKeyWord
	qb.table = table
	return qb
}

func (qb *DefaultSQLQueryBuilder) Set(column string, value any) *DefaultSQLQueryBuilder {
	qb.columns = append(qb.columns, column)
	qb.values = append(qb.values, value)
	return qb
}

func (qb *DefaultSQLQueryBuilder) Delete() *DefaultSQLQueryBuilder {
	qb.query += DeleteKeyWord
	return qb
}

func (qb *DefaultSQLQueryBuilder) Build() (string, []any) {
	switch qb.query {
	case SelectKeyWord:
		return qb.buildSelect()
	case InsertKeyWord:
		return qb.buildInsert()
	case UpdateKeyWord:
		return qb.buildUpdate()
	case DeleteKeyWord:
		return qb.buildDelete()
	default:
		log.Println("Unsupported query type")
		return "", nil
	}
}

func (qb *DefaultSQLQueryBuilder) buildSelect() (string, []any) {
	AddColumnsToSQLQueryString(&qb.query, qb.columns)

	qb.query += " FROM "
	AddTableToSQLQueryString(&qb.query, qb.table)

	if len(qb.conditions) > 0 {
		qb.query += " WHERE "
		qb.param.AddConditions(&qb.query, qb.conditions)
	}

	qb.query += ";"
	return qb.query, qb.args
}

func (qb *DefaultSQLQueryBuilder) buildInsert() (string, []any) {
	qb.query += qb.table

	qb.query += "("
	AddColumnsToSQLQueryString(&qb.query, qb.columns)
	qb.query += ")"

	qb.query += " VALUES ("
	qb.param.AddPlaceholders(&qb.query, qb.columns)
	qb.query += ");"

	return qb.query, qb.values
}

func (qb *DefaultSQLQueryBuilder) buildUpdate() (string, []any) {
	if len(qb.columns) == 0 {
		return qb.query, nil
	}

	qb.query += qb.table + " SET "

	qb.param.AddSetColumns(&qb.query, qb.columns)

	if len(qb.conditions) > 0 {
		qb.query += " WHERE "
		qb.param.AddConditions(&qb.query, qb.conditions)
	}

	qb.query += ";"
	return qb.query, append(qb.values, qb.args...)
}

func (qb *DefaultSQLQueryBuilder) buildDelete() (string, []any) {
	qb.query += "FROM "
	qb.query += qb.table

	qb.param.AddConditions(&qb.query, qb.conditions)

	qb.query += ";"
	return qb.query, qb.args
}
