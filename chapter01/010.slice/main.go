package main

import "fmt"

func main() {
	a := "hello"
	fmt.Println(a)
	aBytes := []byte(a) // byte可以跟string转换
	fmt.Println(aBytes)
	fmt.Println("修改切片内的内容")
	aBytes[0] = 'H'
	a = string(aBytes)
	fmt.Println(a) // Hello

}
