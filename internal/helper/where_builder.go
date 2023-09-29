package helper

import (
	"fmt"
	"reflect"
)

//honestly i preffered raw query rather than using gorm
//this function is used for query builder but its still 50% complete
//where string of raw query is concated with this function result
//example of code
// ================================================================================
// query := `Select * from users %s` the %s is used for concatting the where
// where, value := WhereBuilder("db", entity.User{})
// query := fmt.sprintf(`Select * from users %s`, where)
// ================================================================================

// so the result query looks like this
// Select * from users Where id = ?
// how about the value? we'll use variadic function from library mysql

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
