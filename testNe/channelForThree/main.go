package main

import (
	"fmt"
	"time"
)

var buffer chan string

func init() {
	buffer = make(chan string, 1000)
}
func main() {

	go put()
	go put()
	go put()
	go put()
	go take()
	go take()
	time.Sleep(1 * time.Second)
}

func put() {
	for i := 0; i < 10; i++ {
		buffer <- "11111"
	}
}

func take() {
	for i := 0; i < 20; i++ {
		v := <-buffer
		fmt.Println(v)
	}
}
