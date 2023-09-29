package helper

import (
	"fmt"
	"reflect"
)

func WhereBuilder(tags string, p interface{}) (field []string, value []string) {

	v := reflect.ValueOf(p).Elem()
	t := reflect.TypeOf(p).Elem()

	value = []string{}
	field = []string{}

	for i := 0; i < v.NumField(); i++ {
		value = append(value, v.Field(i).String())
	}

	for i := 0; i < t.NumField(); i++ {
		field = append(field, fmt.Sprintf(`%s = ?`, t.Field(i).Tag.Get(tags)))
	}

	if len(value) == 0 {
		return []string{}, []string{}
	}

	return field, value
}
