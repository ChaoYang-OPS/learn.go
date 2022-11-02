package main

import "fmt"

func (c *Calculator) Add() int {
	fmt.Printf("&c @Calculator %p\n", &c)
	tempResult := c.left + c.right
	c.result = tempResult
	fmt.Println("调用add函数，c的result=", c.result)
	return tempResult
}
func (c Calculator) Sub() int {
	return 0
}

func (c Calculator) Muliple() int {
	return 0
}

func (c Calculator) Divide() int {
	return 0
}

func (c Calculator) Reminder() int {
	return 0
}

func (c *Calculator) SetResult(result int) {
	c.result = result
}
