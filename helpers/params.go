package helpers

import "fmt"

func CheckQueryIsEmpty(query *string) bool {
	return query == nil || *query == ""
}

func CheckColumnsIsEmpty(columns []string) bool {
	return len(columns) == 0
}

func CheckDuplicateColumns(columns []string) bool {
	seen := make(map[string]bool)
	for _, column := range columns {
		if seen[column] {
			return true
		}
		seen[column] = true
	}
	return false
}

func CheckInput(query *string, columns []string) error {
	if CheckQueryIsEmpty(query) {
		return fmt.Errorf("query is empty")
	}

	if CheckColumnsIsEmpty(columns) {
		return fmt.Errorf("columns are empty")
	}

	if CheckDuplicateColumns(columns) {
		return fmt.Errorf("duplicate columns found")
	}
	return nil
}
