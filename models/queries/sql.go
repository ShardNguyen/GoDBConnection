package queries

import (
	"fmt"
	"log"
)

type SQLQueryBuilder struct {
	table      string
	columns    []string
	conditions []string
	args       []any
	query      string
}

func NewSQLQueryBuilder(table string) *SQLQueryBuilder {
	return &SQLQueryBuilder{
		table:      table,
		columns:    []string{},
		conditions: []string{},
		args:       []any{},
		query:      "",
	}
}

func (qb *SQLQueryBuilder) Select(columns ...string) *SQLQueryBuilder {
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
	qb.table = table
	qb.columns = columns
	return qb
}

func (qb *SQLQueryBuilder) Values(values ...any) *SQLQueryBuilder {
	qb.args = values
	return qb
}

func (qb *SQLQueryBuilder) BuildInsert() (string, []any) {
	if len(qb.columns) == 0 {
		log.Println("No columns specified for insert")
		return "", nil
	}

	qb.query = "INSERT INTO " + qb.table + " ("

	qb.query += qb.columns[0]
	for _, column := range qb.columns[1:] {
		qb.query += ", " + column
	}

	qb.query += ") VALUES ("
	qb.query += qb.args[0].(string)
	for _, arg := range qb.args[1:] {
		qb.query += ", " + fmt.Sprintf("%d", arg)
	}

	qb.query += ");"
	return qb.query, qb.args
}

func (qb *SQLQueryBuilder) BuildSelect() (string, []any) {
	qb.query = "SELECT "

	if len(qb.columns) == 0 {
		qb.query += "*"
	}

	if len(qb.columns) > 0 {
		qb.query += qb.columns[0]
		for _, column := range qb.columns[1:] {
			qb.query += ", " + column
		}
	}

	qb.query += " FROM " + qb.table

	if len(qb.conditions) > 0 {
		qb.query += " WHERE " + qb.conditions[0]
		for _, condition := range qb.conditions[1:] {
			qb.query += " AND " + condition
		}
	}

	qb.query += ";"
	return qb.query, qb.args
}
