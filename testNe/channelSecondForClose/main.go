package main

import "fmt"

// chan struct{}  空结构体变量的内存占用为0
func main() {
	c := make(chan int, 2)
	// 往channel中写入数据
	c <- 1
	c <- 2
	close(c) // close can not allow write
	//close(c) // 已经关闭的channl，再次关闭会Panic  panic: close of closed channel
	//c <- 3  // panic: send on closed channel
	// close(nil) 也会Panic
	// for range 之前必须关闭管道,否则会直接死锁
	for item := range c {
		fmt.Println(item)
	}
	//从channel中读取值
	v := <-c // close channel之后，读操作总是会立即返回，
	// 如果channel里面没有元素，则返回零值
	fmt.Println(v) // 0
}
