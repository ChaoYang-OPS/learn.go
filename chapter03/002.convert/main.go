package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	convertType()
	fmt.Println("Finished")
}

func convertType() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("panic--->",e)
			debug.PrintStack()
		}
	}()
	var a interface{}
	a = "string aaa"
	b := a.(int)  // panic: interface conversion: interface {} is string, not int
	fmt.Println(b)
}
