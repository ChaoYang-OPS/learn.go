package main

import "fmt"

func foo(i interface{}) {
	if v, ok := i.(int); ok {
		fmt.Println("is is int", v)
	} else if v, ok := i.(float32); ok {
		fmt.Println("i is is float32", v)
	} else {
		fmt.Println("other type")
	}

}

func f2(i interface{}) {
	switch i.(type) {
	case int:
		v := i.(int)
		fmt.Printf("i is int %d\n", v)
	case float32:
		v := i.(float32)
		fmt.Printf("i is float32 %f\n", v)
	default:
		fmt.Println("other type")
	}
}

func f3(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("i is int %d\n", v)
	case float32:
		fmt.Printf("i is float32 %f\n", v)
	default:
		fmt.Println("other type")
	}
}

func main() {
	foo(123)
	foo(3.2)
	foo(true)
	f2(3)
	var f1 float32
	f2(f1)
	f3(f1)
}
