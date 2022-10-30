package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func Marshal(v interface{}) string {

	value := reflect.ValueOf(v)
	valueType := value.Type()
	switch valueType.Kind() {
	case reflect.Int:
		return strconv.Itoa(int(value.Int()))
	case reflect.String:
		return value.String()
	default:
		return ""
	}
}

func Unmarshal(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return result
}
func main() {
	var v int = 7
	s := Marshal(v)
	fmt.Println(s)
	result := Unmarshal(s)
	fmt.Println(result)
	//json.Marshal()
}
