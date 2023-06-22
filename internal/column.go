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

func validateColumn(sf *reflect.StructField, v *reflect.Value) error {
	if !isRestrictedColumn(sf) {
		return nil
	}

	available := availableColumnValues(sf)
	if !containValue(available, columnValue(v)) {
		return errors.Errorf("field(%v) is invalid", sf.Name)
	}

	return nil
}

func availableColumnValues(sf *reflect.StructField) []string {
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

func columnValue(field *reflect.Value) string {
	return fmt.Sprintf("%v", *field)
}
