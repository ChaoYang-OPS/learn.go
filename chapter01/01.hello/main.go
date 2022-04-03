package main

import "fmt"

func main() {
	var name string = "hello, golang!"
	fmt.Println(name)
	var age int = 35
	fmt.Println(age)
	var tall float64 = 1.70
	fmt.Println(tall)
	// Golang是强类型语言
	// 在赋值过程中，类型必须保持一致

	name = "hello Luckily"
	//age = "Hello"  // '"Hello"' (type string) cannot be represented by the type int 类型不兼容

	//tall = age  // Cannot use 'age' (type int) as the type float64 类型不兼容
	// 变量必须先定义，才能使用

}
