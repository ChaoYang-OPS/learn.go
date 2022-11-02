package main

import "fmt"

func main() {
	var fruit string = "6 apples"
	var watermallan bool = false // or false
	if watermallan {
		fruit = "1 apple"
	} else {
		fruit = "7 apples"
	}
	fmt.Println("buy:", fruit) // buy: 6 apples
}
