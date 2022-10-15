package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string //16 字节
	Age  int    // 8
}

func RefType() {
	typeI := reflect.TypeOf(1)
	typeS := reflect.TypeOf("Hello world")
	typeT := reflect.TypeOf(true)
	fmt.Println(typeI.String())
	fmt.Println(typeS.String())
	fmt.Println(typeI.Kind())
	fmt.Println(typeS.Kind())
	fmt.Println(typeT)
	u1 := User{}
	useType := reflect.TypeOf(u1)
	fmt.Println(useType)
	fmt.Println(useType.Kind() == reflect.Struct) // struct
	fmt.Println(useType.String())
	u2 := &User{}
	use2Type := reflect.TypeOf(u2)
	fmt.Println(use2Type)
	fmt.Println(use2Type.Kind() == reflect.Ptr) //  ptr
	fmt.Println(use2Type.String())

	typeUser3 := use2Type.Elem() // 解析指针
	fmt.Println(typeUser3)
	fmt.Println(typeUser3.Kind() == reflect.Struct) //   struct
	fmt.Println(typeUser3.String())

	fmt.Println(useType.Name())
	fmt.Println(useType.PkgPath())
	fmt.Println(useType.Size()) // 24
}

func main() {
	RefType()
}
