package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := sync.Map{}
	for i := 0; i <100; i++{
		go func(i int) {
			for {
				v,ok := m.Load(i)
				if !ok{
					continue
				}
				m.Store(i,v.(int)+1)
				fmt.Println("i=",v)
			}
		}(i)
	}
	time.Sleep(10 * time.Second)
}
