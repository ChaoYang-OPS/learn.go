package main

import (
	"fmt"
	"sync"
	"time"
)

func AddTimePrint() {
	fmt.Println(time.Now())
}
func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		AddTimePrint()
	}()
	go func() {
		defer wg.Done()
		AddTimePrint()
	}()
	wg.Wait()
	fmt.Println(".....")
}
