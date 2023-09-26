package dbmap

import (
	"reflect"
)

func Columns(t reflect.Type) []string {
	vs := make([]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		vs[i] = t.Field(i).Tag.Get("db")
	}
	return vs
}

func Values(s any) []any {
	v := reflect.ValueOf(s)
	vs := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		vs[i] = v.Field(i).Interface()
	}
	return vs
}

func Scan(s any) []any {
	vo := reflect.ValueOf(s).Elem()
	vs := make([]interface{}, vo.NumField())
	for i := 0; i < vo.NumField(); i++ {
		f := vo.Field(i)
		vs[i] = f.Addr().Interface()
	}
	return vs
}
