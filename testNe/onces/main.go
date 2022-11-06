package main

import (
	"fmt"
	"sync"
	"time"
)

var onces sync.Once

var a int = 10

func main() {
	go func() {
		onces.Do(func() {
			fmt.Println("First")
			a++
		})
	}()
	go func() {
		onces.Do(func() {
			fmt.Println("Second")
			a++
		})
	}()
	time.Sleep(2 * time.Second)
	fmt.Println(a)
}
