package internal

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

const (
	columnKey = "column"

	columEnumType = "enum"
)

func RowValues(v any) []string {
	values := reflect.ValueOf(v).Elem()
	var ret []string
	for i := 0; i < values.NumField(); i++ {
		field := values.Field(i)
		ret = append(ret, stringValue(&field))
	}
	return ret
}

func stringValue(field *reflect.Value) string {
	return fmt.Sprintf("%v", *field)
}

func validateColumn(sf *reflect.StructField, v *reflect.Value) error {
	if !isRestrictedColumn(sf) {
		return nil
	}

	available := availableValues(sf)
	if !containValue(available, stringValue(v)) {
		return errors.Errorf("field(%v) is invalid", sf.Name)
	}

	return nil
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

func containValue(available []string, value string) bool {
	for _, item := range available {
		if item == value {
			return true
		}
	}
	return false
}

func availableValues(sf *reflect.StructField) []string {
	value, ok := sf.Tag.Lookup(columnKey)
	if !ok {
		return nil
	}

	if value != columEnumType {
		return nil
	}

	enum := sf.Tag.Get("enum")
	return strings.Split(enum, ",")
}

func isRestrictedColumn(sf *reflect.StructField) bool {
	value, ok := sf.Tag.Lookup(columnKey)
	if !ok {
		return false
	}

	if value != columEnumType {
		return false
	}

	return true
}
