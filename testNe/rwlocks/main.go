package main

import (
	"fmt"
	"sync"
)

var lock sync.RWMutex

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(3)
	go func() {
		defer wg.Done()
		lock.RLock()
		fmt.Println("A lock success")
		//lock.Unlock()
	}()
	go func() {
		defer wg.Done()
		lock.RLock()
		fmt.Println("B lock success")
	}()
	go func() {
		defer wg.Done()
		lock.RLock()
		fmt.Println("C lock success")
	}()
	wg.Wait()
	fmt.Println("......")
}
