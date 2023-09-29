package helper

import "fmt"

func InterfaceToInt(v interface{}) int {
	return v.(int)
}

func InterfaceToString(v interface{}) string {
	return fmt.Sprintf(`%v`, v)
}
