package main

import (
	"fmt"
	"time"
)

func main1() {
	c := make(chan int)
	go func(ch chan<- int, x int) {
		time.Sleep(time.Second)
		ch <- x * x
	}(c, 3)
	done := make(chan struct{})
	go func(ch <-chan int) {
		n := <-ch
		fmt.Println(n)
		time.Sleep(time.Second)
		done <- struct{}{}
	}(c)
	<-done
	fmt.Println("bye")
}

func main2() {
	c := make(chan int, 2)
	c <- 3
	c <- 5
	close(c)
	fmt.Println(len(c), cap(c)) // 2 2
	x, ok := <-c
	fmt.Println(x, ok)          //  3 true
	fmt.Println(len(c), cap(c)) // 1 2
	x, ok = <-c
	fmt.Println(x, ok)          // 5 true
	fmt.Println(len(c), cap(c)) // 0  2
	x, ok = <-c
	fmt.Println(x, ok) // 0 false
}

func main3() {
	var ball = make(chan string)
	killBall := func(playerName string) {
		for {
			fmt.Println(<-ball, "传球")
			time.Sleep(time.Second)
			ball <- playerName
		}
	}
	go killBall("张三")
	go killBall("李四")
	go killBall("王二麻子")
	go killBall("刘大")
	ball <- "裁判"    // 开球
	var c chan bool // 一个零值nil通道
	<-c             // 永久阻塞在此
}

func main4() {
	fibonacci := func() chan uint64 {
		c := make(chan uint64)
		go func() {
			var x, y uint64 = 0, 1
			for ; y < (1 << 63); c <- y {
				x, y = y, x+y
			}
			close(c)
		}()
		return c
	}
	c := fibonacci()
	//for x, ok := <-c; ok; x, ok = <-c {
	//	time.Sleep(time.Second)
	//	fmt.Println(x)
	//}
	for x := range c {
		time.Sleep(time.Second)
		fmt.Println(x)
	}
}
