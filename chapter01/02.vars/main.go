package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	var hello string = "Hello golang!"
	var world = "world"
	fmt.Println(hello, world)
	//小数 := 1.234
	//fmt.Println(小数) // 1.234
	var int3, int4 uint = 33, 44
	fmt.Println(int3, int4)
	var (
		int6, int7 = 6, 7
	)
	fmt.Println(int6, int7)
	// 变量必须先定义后使用，且必须被用到
	// Golang会为每个变量设置默认值
	var ho, ver float64 = 3, 4.123
	var sc = ho * ver
	fmt.Println(ho * ver)
	fmt.Println(sc)
	// 变量不能重名
	// Golang会根据类型推断
	var newname string
	fmt.Println("newname=", newname)

	// 类型推断
	a1 := "hello"
	fmt.Println(reflect.TypeOf(a1)) // string
	a2 := 3
	fmt.Println(reflect.TypeOf(a2)) //int
	a3 := 3.0
	fmt.Println(reflect.TypeOf(a3)) //float64
	a4 := &a3
	fmt.Println(reflect.TypeOf(a4)) //*float64
	// 变量类型转换及其代价
	// 1. 每种数据类型都有大小
	// 2. 每种数据类型的存储方式都不同
	// 3. 不是每种数据类型之间者可以转换的
	// 4. 可转换数据类型之间的转换是有代价的
	var f1 float64 = 1.234
	var i1 int = int(f1)
	fmt.Println("f1:", f1, "i1:", i1) // f1: 1.234 i1: 1
	var a6 uint = math.MaxUint64
	fmt.Println(a6)
	var a7 int = int(a6)
	fmt.Println(a6, a7) // 18446744073709551615 -1

	//var nameOfSquare string
}
