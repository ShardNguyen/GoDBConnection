package main

import (
	"GoDBConnection/models/queries"
	"fmt"
)

func main() {
	// Example usage of SQLQueryBuilder
	qb := queries.NewSQLQueryBuilder("users")
	// query, args := qb.Select("name").Where("age > ?", 18).Where("name > ?", "Smith").BuildSelect()
	query, args := qb.InsertInto("bye", "name", "age").Values("John Doe", 30).BuildInsert()

	fmt.Println(query)
	fmt.Println(args)
}
