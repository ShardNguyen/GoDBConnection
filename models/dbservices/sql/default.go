package sqldbsv

import (
	"GoDBConnection/helpers"
	sqlqueries "GoDBConnection/models/queries/sql"
	"fmt"
)

type DefaultSQLDatabaseService struct {
	qb *sqlqueries.DefaultSQLQueryBuilder
}

func NewDefaultSQLDatabaseService() *DefaultSQLDatabaseService {
	return &DefaultSQLDatabaseService{
		qb: sqlqueries.NewSQLQueryBuilder(),
	}
}

func (dbsv *DefaultSQLDatabaseService) CreateTable() error {
	return nil
}

func (dbsv *DefaultSQLDatabaseService) DropTable() error {
	return nil
}

func (dbsv *DefaultSQLDatabaseService) InsertRow(tableName string, entity any) error {
	columns, values, err := helpers.GetStructAttNamesAndValues(entity)

	if err != nil {
		return err
	}

	query, args, err := dbsv.qb.InsertInto(tableName, columns...).Values(values...).Build()

	if err != nil {
		return err
	}

	fmt.Println(*query)
	fmt.Println(args)

	return nil
}

func (dbsv *DefaultSQLDatabaseService) UpdateRow(tableName string, entity any) error {

	// query, args, err := dbsv.qb.Update("users").Set("name", "John Doe").Set("age", 30).Where("id = ", 1).Where("id = ", 2).Build()
	return nil
}

func (dbsv *DefaultSQLDatabaseService) DeleteRow(tableName string, conditions []helpers.Condition) error {
	dbsv.qb = dbsv.qb.Delete().From(tableName)

	if len(conditions) > 0 {
		for _, condition := range conditions {
			dbsv.qb = dbsv.qb.Where(fmt.Sprint(condition.Column, condition.Operator), condition.Value)
		}
	}

	query, args, err := dbsv.qb.Build()
	if err != nil {
		return err
	}

	fmt.Println(*query)
	fmt.Println(args)

	return nil
}

func (dbsv *DefaultSQLDatabaseService) SelectRow(tableName string, columns []string, conditions []helpers.Condition) error {

	dbsv.qb = dbsv.qb.Select(columns...).From(tableName)

	if len(conditions) > 0 {
		for _, condition := range conditions {
			dbsv.qb = dbsv.qb.Where(fmt.Sprint(condition.Column, condition.Operator), condition.Value)
		}
	}

	query, args, err := dbsv.qb.Build()

	if err != nil {
		return err
	}

	fmt.Println(*query)
	fmt.Println(args)

	return nil
}

func (dbsv *DefaultSQLDatabaseService) SelectAllRows(tableName string, columns []string) error {
	query, args, err := dbsv.qb.Select(columns...).From(tableName).Build()

	if err != nil {
		return err
	}

	fmt.Println(*query)
	fmt.Println(args)

	return nil
}
