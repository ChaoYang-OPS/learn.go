package main

import (
	"fmt"
	"sync"
)

func F() {
	defer fmt.Println("currently Defer in F")
	defer fmt.Println("currently Defer in F1")
	defer fmt.Println("currently Defer in F2")
	fmt.Println("Hello word")
	defer func() {
		recover()
	}()
	panic("ERRORS For panic IN F")
	defer fmt.Println("currently Defer in F3")
	defer fmt.Println("currently Defer in F4")
}

//Hello word
//currently Defer in F2
//currently Defer in F1
//currently Defer in F
//panic: ERRORS For panic IN F

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		F()
	}()
	wg.Wait()
}
