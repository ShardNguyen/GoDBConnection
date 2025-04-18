package helpers

type Condition struct {
	Column   string
	Operator string
	Value    any
}

func NewCondition(column, operator string, value any) *Condition {
	return &Condition{
		Column:   column,
		Operator: operator,
		Value:    value,
	}
}
