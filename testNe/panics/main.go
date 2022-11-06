package main

import (
	"fmt"
	"sync"
)

func F() {
	defer fmt.Println("currently Defer in F")
	defer fmt.Println("currently Defer in F1")
	defer fmt.Println("currently Defer in F2")
	fmt.Println("Hello world")
	defer func() {
		recover() // 从panic结束本协程，但没有结束整个进程
	}()
	panic("ERRORS For panic IN F")
	defer fmt.Println("currently Defer in F3")
	defer fmt.Println("currently Defer in F4")
}

//Hello world
//currently Defer in F2
//currently Defer in F1
//currently Defer in F
//panic: ERRORS For panic IN F

func main() {
	//wg := sync.WaitGroup{}
	wgp := new(sync.WaitGroup)
	wgp.Add(1)
	go func() {
		defer wgp.Done()
		F()
	}()
	wgp.Wait()
	fmt.Println(".....ED")
}
