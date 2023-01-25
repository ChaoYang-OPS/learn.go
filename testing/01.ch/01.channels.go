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

func main5() {
	var c chan struct{} // nil
	select {
	case <-c:
	case c <- struct{}{}:
	default:
		fmt.Println("GO Here")
	}
}

func main() {
	c := make(chan string, 2)
	trySend := func(v string) {
		select {
		case c <- v:
		default: // 如果C的缓冲已满，则执行默认分支

		}
	}
	tryReceive := func() string {
		select {
		case v := <-c:
			return v
		default:
			return "-" // 如果C的缓冲为空，则执行默认分支
		}
	}
	trySend("Hello")          // 发送成功
	trySend("Hi")             // 发送成功
	trySend("Bye")            // 发送失败，但不会阻塞
	fmt.Println(tryReceive()) // Hello
	fmt.Println(tryReceive()) //Hi
	// 下面这个行将接收失败
	fmt.Println(tryReceive()) // -
}
