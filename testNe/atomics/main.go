package main

import (
	"fmt"
	"sync"
)

var n int32
var lock sync.RWMutex

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			//n++
			//atomic.AddInt32(&n, 1) 也可以替换Lock
			lock.Lock() // 串行执行
			n++
			lock.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(n) // 1000
}
