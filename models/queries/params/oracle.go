package params

import "fmt"

type OracleParams struct {
}

func NewOracleQueryParams() *OracleParams {
	return &OracleParams{}
}

func (o *OracleParams) AddPlaceholders(query *string, columns []string) {
	if len(columns) == 0 {
		return
	}

	*query += ":" + fmt.Sprint(columns[0])

	for i := 1; i < len(columns); i++ {
		*query += ", :" + fmt.Sprint(columns[1])
	}
}

func (o *OracleParams) AddSetColumns(query *string, columns []string) {
	if len(columns) == 0 {
		return
	}

	*query += columns[0] + " = :" + fmt.Sprint(columns[0])

	for i := 1; i < len(columns); i++ {
		*query += ", " + columns[i] + " = :" + fmt.Sprint(columns[i])
	}
}

func (o *OracleParams) AddConditions(query *string, conditions []string) {
	if len(conditions) == 0 {
		return
	}

	*query += conditions[0] + " :" + fmt.Sprint(conditions[0])

	for i := 1; i < len(conditions); i++ {
		*query += ", " + conditions[i] + " :" + fmt.Sprint(conditions[i])
	}
}
