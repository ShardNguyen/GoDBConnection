package queries

import (
	"log"
)

const (
	SelectKeyWord = "SELECT"
	InsertKeyWord = "INSERT INTO"
	UpdateKeyWord = "UPDATE"
	DeleteKeyWord = "DELETE "
)

type SQLQueryBuilder struct {
	table      string
	columns    []string
	conditions []string
	args       []any
	values     []any
	query      string
}

func NewSQLQueryBuilder(table string) *SQLQueryBuilder {
	return &SQLQueryBuilder{
		table:      table,
		columns:    []string{},
		conditions: []string{},
		args:       []any{},
		values:     []any{},
		query:      "",
	}
}

func (qb *SQLQueryBuilder) Select(columns ...string) *SQLQueryBuilder {
	qb.query += SelectKeyWord
	qb.columns = columns
	return qb
}

func (qb *SQLQueryBuilder) From(table string) *SQLQueryBuilder {
	qb.table = table
	return qb
}

func (qb *SQLQueryBuilder) Where(condition string, arg any) *SQLQueryBuilder {
	qb.conditions = append(qb.conditions, condition)
	qb.args = append(qb.args, arg)
	return qb
}

func (qb *SQLQueryBuilder) InsertInto(table string, columns ...string) *SQLQueryBuilder {
	qb.query += InsertKeyWord
	qb.table = table
	qb.columns = columns
	return qb
}

func (qb *SQLQueryBuilder) Values(values ...any) *SQLQueryBuilder {
	qb.values = values
	return qb
}

func (qb *SQLQueryBuilder) Update(table string) *SQLQueryBuilder {
	qb.query += UpdateKeyWord
	qb.table = table
	return qb
}

func (qb *SQLQueryBuilder) Set(column string, value any) *SQLQueryBuilder {
	qb.columns = append(qb.columns, column)
	qb.values = append(qb.values, value)
	return qb
}

func (qb *SQLQueryBuilder) Delete() *SQLQueryBuilder {
	qb.query += DeleteKeyWord
	return qb
}

func (qb *SQLQueryBuilder) Build() (string, []any) {
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

func (qb *SQLQueryBuilder) buildSelect() (string, []any) {
	if len(qb.columns) == 0 {
		qb.query += "*"
	}

	AddColumnsToSQLQueryString(&qb.query, qb.columns)

	qb.query += " FROM " + qb.table

	AddConditionsToSQLQueryString(&qb.query, qb.conditions)

	qb.query += ";"
	return qb.query, qb.args
}

func (qb *SQLQueryBuilder) buildInsert() (string, []any) {
	qb.query += qb.table

	qb.query += "("
	AddColumnsToSQLQueryString(&qb.query, qb.columns)
	qb.query += ")"

	qb.query += " VALUES ("
	qb.query += "?"

	for range qb.values[1:] {
		qb.query += ", " + "?"
	}

	qb.query += ");"

	return qb.query, qb.values
}

func (qb *SQLQueryBuilder) buildUpdate() (string, []any) {
	qb.query += qb.table + " SET "

	if len(qb.columns) == 0 {
		return "", nil
	}

	qb.query += qb.columns[0] + " = ?"
	for _, col := range qb.columns[1:] {
		qb.query += ", " + col + " = ?"
	}

	AddConditionsToSQLQueryString(&qb.query, qb.conditions)

	qb.query += ";"
	return qb.query, append(qb.values, qb.args...)
}

func (qb *SQLQueryBuilder) buildDelete() (string, []any) {
	qb.query += "FROM "
	qb.query += qb.table

	AddConditionsToSQLQueryString(&qb.query, qb.conditions)

	qb.query += ";"
	return qb.query, qb.args
}
