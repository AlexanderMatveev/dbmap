package dbmap

import (
	"reflect"
)

const tagName = "db"

func Columns(t reflect.Type) []string {
	var vs []string
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get(tagName)
		if tag == "" {
			continue
		}
		vs = append(vs, tag)
	}
	return vs
}

func Values(s any) []any {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)
	var vs []any
	for i := 0; i < v.NumField(); i++ {
		if t.Field(i).Tag.Get(tagName) == "" {
			continue
		}
		vs = append(vs, v.Field(i).Interface())
	}
	return vs
}

func Scan(s any) []any {
	v := reflect.ValueOf(s).Elem()
	t := reflect.TypeOf(s).Elem()
	var vs []any
	for i := 0; i < v.NumField(); i++ {
		if t.Field(i).Tag.Get(tagName) == "" {
			continue
		}
		vs = append(vs, v.Field(i).Addr().Interface())
	}
	return vs
}
