package queries

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
	qb.query += "SELECT "
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
	qb.query += "INSERT INTO "
	qb.table = table
	qb.columns = columns
	return qb
}

func (qb *SQLQueryBuilder) Values(values ...any) *SQLQueryBuilder {
	qb.values = values
	return qb
}

func (qb *SQLQueryBuilder) Update(table string) *SQLQueryBuilder {
	qb.query += "UPDATE "
	qb.table = table
	return qb
}

func (qb *SQLQueryBuilder) Set(column string, value any) *SQLQueryBuilder {
	qb.columns = append(qb.columns, column)
	qb.values = append(qb.values, value)
	return qb
}

func (qb *SQLQueryBuilder) Build() (string, []any) {
	switch qb.query {
	case "SELECT ":
		return qb.buildSelect()
	case "INSERT INTO ":
		return qb.buildInsert()
	case "UPDATE ":
		return qb.buildUpdate()
	default:
		return "", nil
	}
}

func (qb *SQLQueryBuilder) buildSelect() (string, []any) {
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

func (qb *SQLQueryBuilder) buildInsert() (string, []any) {
	qb.query += qb.table

	if len(qb.columns) != 0 {
		qb.query += "("
		qb.query += qb.columns[0]
		for _, column := range qb.columns[1:] {
			qb.query += ", " + column
		}
		qb.query += ")"
	}

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

	if len(qb.conditions) > 0 {
		qb.query += " WHERE " + qb.conditions[0]

		for _, condition := range qb.conditions[1:] {
			qb.query += " AND " + condition
		}
	}

	qb.query += ";"
	return qb.query, append(qb.values, qb.args...)
}
