package main

import (
	sqldbsv "GoDBConnection/models/dbservices/sql"
)

func main() {
	/*
		errorLogger := logs.NewErrorLogger()

		// Example usage of SQLQueryBuilder
		qb := queries.NewPostgresSQLQueryBuilder("users")
		// query, args := qb.Select().Where("age > ", 18).Where("name = ", "Smith").Build()
		// query, args := qb.InsertInto("bye", "name", "age").Values("John Doe", 30).Build()
		query, args, err := qb.Update("users").Set("name", "John Doe").Set("age", 30).Where("id = ", 1).Where("id = ", 2).Build()
		// query, args := qb.Delete().Where("id = ?", 1).Build()

		if err != nil {
			errorLogger.ConsoleLogError(err)
			return
		}

		fmt.Println(*query)
		fmt.Println(args)
	*/

	dbsv := sqldbsv.NewDefaultSQLDatabaseService()

	/*
		user := entities.User{
			ID:    1,
			Name:  "john_doe",
			Email: "john@example.com",
		}

		dbsv.InsertRow("users", user)
	*/

	dbsv.SelectAllRows("users", []string{"name", "email", "trolley"})
}
