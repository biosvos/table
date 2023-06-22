package internal

import (
	"github.com/pkg/errors"
	"reflect"
)

func RowValues(v any) []string {
	values := reflect.ValueOf(v).Elem()
	var ret []string
	for i := 0; i < values.NumField(); i++ {
		field := values.Field(i)
		ret = append(ret, columnValue(&field))
	}
	return ret
}

func ValidateRow(v any) error {
	types := reflect.TypeOf(v).Elem()
	values := reflect.ValueOf(v).Elem()
	for i := 0; i < types.NumField(); i++ {
		field := types.Field(i)
		value := values.Field(i)
		err := validateColumn(&field, &value)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}
