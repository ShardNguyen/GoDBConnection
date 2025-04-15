package main

import (
	"GoDBConnection/models/queries"
	"fmt"
)

func main() {
	// Example usage of SQLQueryBuilder
	qb := queries.NewSQLQueryBuilder("users")
	//query, args := qb.Select("name").Where("age > ?", 18).Where("name = ?", "Smith").Build()
	// query, args := qb.InsertInto("bye", "name", "age").Values("John Doe", 30).Build()
	// query, args := qb.Update("users").Set("name", "John Doe").Set("age", 30).Where("id = ?", 1).Build()
	query, args := qb.Delete().Where("id = ?", 1).Build()

	fmt.Println(query)
	fmt.Println(args)
}
