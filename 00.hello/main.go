package main

import "fmt"

func main() {
	// GOOS=windows GOARCH=amd64 go build ./main.go
	// GOOS=linux GOARCH=amd64 go build ./main.go
	fmt.Println("您好, golang!")
	fmt.Println("1.打开冰箱")
	fmt.Println("2.把大象装进去")
	fmt.Println("3.关闭冰箱")
	//
	fmt.Println("你一定要随机应变，千万不可投机取巧")
	fmt.Println("忍人所不忍，能人所不能")
	fmt.Println("做你认为正确的事")
	fmt.Println("")
}