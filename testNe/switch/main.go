package main

import "fmt"

func basic() {
	color := "yellow"
	switch color {
	case "red": // if color == "red"
		fmt.Println("color is red")
	case "green": // else if color =="green"
		fmt.Println("color is green")
	case "black": // else if color =="black"
		fmt.Println("color is black")
	default: // else
		fmt.Println("unknown")
	}
}

func fall(age int) {
	switch {
	case age > 50:
		fmt.Println("当总统")
		fallthrough
	case age > 25:
		fmt.Println("生孩子")
		fallthrough
	case age > 22:
		fmt.Println("结婚")
		fallthrough
	case age > 18:
		fmt.Println("开车")
	}
}
func add(n int) int {
	return n + 1
}

func pos(n int) bool {
	return n > 0
}
func expression() {
	var a int = 10
	const B = 13
	switch add(a) {
	case 7, 8, 11:
		fmt.Println("sum is 7 or 8 or 11")
	case 5 + 9:
		fmt.Println("sum is 5")
	case sub(12):
		fmt.Println("sum is 12")
	case B:
		fmt.Println("sum is 11")
	default:
		fmt.Println("ERROR")
	}
}
func switchCondition() {
	switch {
	case 5 > 8:
		fmt.Println("5>8")
	case add(5) == 9:
		fmt.Println("add(5)==9")
	case add(7) == 9:
		fmt.Println("....")
	case pos(5):
		fmt.Println(".....1")
	}
}
func sub(n int) int {
	return n - 1
}
func main() {
	basic()
	expression()
	switchCondition()
	fall(60)
}
