package main

import (
	"github.com/biosvos/go-template/internal"
	"log"
	"reflect"
	"strings"
	"testing"
)

func TestS(t *testing.T) {
	type Issue struct {
		Id       uint64
		Title    string
		Kind     string `column:"enum" enum:"a,b,c,d"`
		Context  string
		Contents string
		State    string `column:"enum" enum:"a,b,c,d"`
	}
	issues := []*Issue{
		{
			Id:       1,
			Title:    "a",
			Kind:     "b",
			Context:  "c",
			Contents: "d",
			State:    "e",
		},
		{
			Id:       2,
			Title:    "f",
			Kind:     "g",
			Context:  "h",
			Contents: "i",
			State:    "j",
		},
	}

	for _, issue := range issues {
		values := internal.RowValues(issue)
		log.Println(values)
	}
}

func TestName(t *testing.T) {
	type Issue struct {
		Title    string
		Kind     string `column:"enum" enum:"a,b,c,d"`
		Context  string
		Contents string
		State    string `column:"enum" enum:"a,b,c,d"`
	}
	issue := Issue{
		Title:    "a",
		Kind:     "b",
		Context:  "c",
		Contents: "d",
		State:    "e",
	}
	elem := reflect.TypeOf(&issue).Elem()
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)

		value, ok := field.Tag.Lookup("column")
		t.Logf("%+v %+v", value, ok)
		if ok && value == "enum" {
			value, ok := field.Tag.Lookup("enum")
			if ok {
				split := strings.Split(value, ",")
				t.Logf("%+v", split)
			}
		}
	}
}
