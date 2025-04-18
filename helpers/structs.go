package helpers

import (
	"errors"
	"reflect"
)

func GetStructAttNamesAndValues(obj any) ([]string, []any, error) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	// Ensure we are working with a struct
	if t.Kind() != reflect.Struct {
		return nil, nil, errors.New("provided value is not a struct")
	}

	// Iterate over struct fields
	names := []string{}
	values := []any{}

	for i := range t.NumField() {
		field := t.Field(i)
		value := v.Field(i)

		names = append(names, field.Name)
		values = append(values, value.Interface())
	}

	return names, values, nil
}
