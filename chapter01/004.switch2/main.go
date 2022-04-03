package main

import (
	"fmt"
	"reflect"
)

func main() {
	var money interface{} = 10
	switch newMoney := money.(type) {
	case int:
		fmt.Println("money is int")
		fmt.Println(newMoney + 3)
	case int64:
		fmt.Println("money is int64")
	case float64:
		fmt.Println("money is float64")
	default:
		fmt.Println("money is unknow")
	}
	fmt.Println(money, reflect.TypeOf(money)) // 10 string
	money = 3
	fmt.Println(money, reflect.TypeOf(money)) // 3 int
}
