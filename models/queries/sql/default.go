package sqlqueries

import (
	"errors"

	"GoDBConnection/helpers"
	"GoDBConnection/models/queries/sql/params"
)

const (
	SelectKeyWord = "SELECT "
	InsertKeyWord = "INSERT INTO "
	UpdateKeyWord = "UPDATE "
	DeleteKeyWord = "DELETE "
	CreateKeyWord = "CREATE "
	DropKeyWord   = "DROP "
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

func NewSQLQueryBuilder() *DefaultSQLQueryBuilder {
	return &DefaultSQLQueryBuilder{
		table:      "",
		columns:    []string{},
		conditions: []string{},
		args:       []any{},
		values:     []any{},
		query:      "",
		param:      params.NewDefaultQueryParams(),
	}
}

// Select initializes the SQL query with the SELECT keyword and sets the columns to be selected.
func (qb *DefaultSQLQueryBuilder) Select(columns ...string) *DefaultSQLQueryBuilder {
	qb.query += SelectKeyWord
	qb.columns = columns
	return qb
}

// From sets the table from which to select data.
func (qb *DefaultSQLQueryBuilder) From(table string) *DefaultSQLQueryBuilder {
	qb.table = table
	return qb
}

// Where adds a condition to the SQL query.
func (qb *DefaultSQLQueryBuilder) Where(condition string, arg any) *DefaultSQLQueryBuilder {
	qb.conditions = append(qb.conditions, condition)
	qb.args = append(qb.args, arg)
	return qb
}

// InsertInto initializes the SQL query with the INSERT INTO keyword and sets the table and columns for insertion.
func (qb *DefaultSQLQueryBuilder) InsertInto(table string, columns ...string) *DefaultSQLQueryBuilder {
	qb.query += InsertKeyWord
	qb.table = table
	qb.columns = columns
	return qb
}

// Values sets the values to be inserted into the table.
func (qb *DefaultSQLQueryBuilder) Values(values ...any) *DefaultSQLQueryBuilder {
	qb.values = values
	return qb
}

// Update initializes the SQL query with the UPDATE keyword and sets the table to be updated.
func (qb *DefaultSQLQueryBuilder) Update(table string) *DefaultSQLQueryBuilder {
	qb.query += UpdateKeyWord
	qb.table = table
	return qb
}

// Set adds a column and its corresponding value to be updated in the SQL query.
func (qb *DefaultSQLQueryBuilder) Set(column string, value any) *DefaultSQLQueryBuilder {
	qb.columns = append(qb.columns, column)
	qb.values = append(qb.values, value)
	return qb
}

// Delete initializes the SQL query with the DELETE keyword.
func (qb *DefaultSQLQueryBuilder) Delete() *DefaultSQLQueryBuilder {
	qb.query += DeleteKeyWord
	return qb
}

func (qb *DefaultSQLQueryBuilder) Create() *DefaultSQLQueryBuilder {
	qb.query += CreateKeyWord
	return qb
}

func (qb *DefaultSQLQueryBuilder) Drop() *DefaultSQLQueryBuilder {
	qb.query += DropKeyWord
	return qb
}

func (qb *DefaultSQLQueryBuilder) Table(table string) *DefaultSQLQueryBuilder {
	qb.table = table
	return qb
}

// Build constructs the SQL query and its arguments based on the initialized parameters.
func (qb *DefaultSQLQueryBuilder) Build() (*string, []any, error) {
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
		return nil, nil, errors.New("invalid SQL query type")
	}
}

func (qb *DefaultSQLQueryBuilder) buildSelect() (*string, []any, error) {
	AddColumns(&qb.query, qb.columns)

	qb.query += " FROM "
	if err := AddTable(&qb.query, qb.table); err != nil {
		return nil, nil, err
	}

	if len(qb.conditions) > 0 {
		qb.query += " WHERE "

		if err := qb.param.AddConditions(&qb.query, qb.conditions); err != nil {
			return nil, nil, err
		}
	}

	qb.query += ";"
	return &qb.query, qb.args, nil
}

func (qb *DefaultSQLQueryBuilder) buildInsert() (*string, []any, error) {
	qb.query += qb.table

	qb.query += "("
	AddColumns(&qb.query, qb.columns)
	qb.query += ")"

	qb.query += " VALUES ("
	if err := qb.param.AddPlaceholders(&qb.query, qb.columns); err != nil {
		return nil, nil, err
	}
	qb.query += ");"

	return &qb.query, qb.values, nil
}

func (qb *DefaultSQLQueryBuilder) buildUpdate() (*string, []any, error) {
	if helpers.CheckColumnsIsEmpty(qb.columns) {
		return nil, nil, errors.New("no columns to update")
	}

	qb.query += qb.table + " SET "

	qb.param.AddSetColumns(&qb.query, qb.columns)

	if len(qb.conditions) > 0 {
		qb.query += " WHERE "
		if err := qb.param.AddConditions(&qb.query, qb.conditions); err != nil {
			return nil, nil, err
		}
	}

	qb.query += ";"
	return &qb.query, append(qb.values, qb.args...), nil
}

func (qb *DefaultSQLQueryBuilder) buildDelete() (*string, []any, error) {
	qb.query += "FROM "
	qb.query += qb.table

	if !helpers.CheckColumnsIsEmpty(qb.conditions) {
		qb.query += " WHERE "
		if err := qb.param.AddConditions(&qb.query, qb.conditions); err != nil {
			return nil, nil, err
		}
	}

	qb.query += ";"
	return &qb.query, qb.args, nil
}
