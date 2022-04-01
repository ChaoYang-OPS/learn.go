package main

import "fmt"

func main() {
	var a, b int = 30, 11
	fmt.Println("a+b", a+b)
	fmt.Println("a-b", a-b)
	fmt.Println("a*b", a*b)
	fmt.Println("a/b", a/b)
	fmt.Println("a%b", a%b)
	fmt.Println(true && false)          // false
	fmt.Println(true && false == false) // true
	fmt.Println("a>b=", a > b)          // a>b= true
	fmt.Println("a<b=", a < b)          // a<b= false
}
