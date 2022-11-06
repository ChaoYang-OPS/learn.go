package main

import (
	"fmt"
	"time"
)

func main() {
	//  struct{}
	ch := make(chan struct{}, 1)
	ch <- struct{}{}
	go func() {
		time.Sleep(3 * time.Second)
		<-ch
		fmt.Println("sub routerline 1 over")
	}()
	ch <- struct{}{}
	fmt.Println("send to channle in main routerline")

	go func() {
		time.Sleep(2 * time.Second)
		ch <- struct{}{}
		fmt.Println("sub route 2 over")
	}()
	time.Sleep(3 * time.Second)
	<-ch
	fmt.Println("main exit...")
	time.Sleep(1 * time.Second)
}
