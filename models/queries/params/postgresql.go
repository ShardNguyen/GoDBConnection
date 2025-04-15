package params

import (
	"fmt"
)

type PostgresSQLQueryParams struct {
	counter int
}

func NewPostgresSQLQueryParams() *PostgresSQLQueryParams {
	return &PostgresSQLQueryParams{
		counter: 1,
	}
}

func (p *PostgresSQLQueryParams) AddPlaceholders(query *string, columns []string) {
	*query += "$" + fmt.Sprint(p.counter)
	p.counter++

	for i := 2; i <= len(columns); i++ {
		*query += ", $" + fmt.Sprint(p.counter)
		p.counter++
	}
}

func (p *PostgresSQLQueryParams) AddSetColumns(query *string, columns []string) {
	if len(columns) == 0 {
		return
	}

	*query += columns[0] + " = $" + fmt.Sprint(p.counter)
	p.counter++

	for i := 1; i < len(columns); i++ {
		*query += ", " + columns[i] + " = $" + fmt.Sprint(p.counter)
		p.counter++
	}
}

func (p *PostgresSQLQueryParams) AddConditions(query *string, conditions []string) {
	if len(conditions) == 0 {
		return
	}

	*query += conditions[0] + "$" + fmt.Sprint(p.counter)
	p.counter++

	for i := 1; i < len(conditions); i++ {
		*query += ", " + conditions[i] + " = $" + fmt.Sprint(p.counter)
		p.counter++
	}
}

func (p *PostgresSQLQueryParams) ResetCounter() {
	p.counter = 1
}
