package main

import (
	"encoding/json"
	"errors"
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
　
type User struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age"`
}

func Unmarshal(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return result
}

func UnMarshal(v interface{}, data []byte) error {
	value := reflect.ValueOf(v)
	typ := value.Type()
	if typ.Kind() != reflect.Ptr {
		return errors.New("v must be pointer")
	}
	s := string(data)
	typ = typ.Elem()
	value = value.Elem()
	switch typ.Kind() {
	case reflect.Int:
		if i, err := strconv.ParseInt(s, 10, 64); err == nil {
			value.SetInt(i)
		} else {
			return err
		}
	case reflect.String:
		value.SetString(s)
	case reflect.Bool:
		// string to bool
		if b, err := strconv.ParseBool(s); err != nil {
			return err
		} else {
			value.SetBool(b)
		}
	case reflect.Float64:
	default:
		return errors.New("unsupported this data")
	}
	return nil
}
func main() {
	var v int = 7
	s := Marshal(v)
	fmt.Println(s)
	result := Unmarshal(s)
	fmt.Println(result)
	//json.Marshal()
	u := User{
		Name: "TF",
		Age:  19,
	}
	object, err := json.Marshal(u)
	if err != nil {
		return
	}
	fmt.Println(string(object))
}
