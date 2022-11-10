package main

import (
	"fmt"
	"os"
	"time"
)

// 倒计时

func countDown(countCh chan int, n int, finishCh chan struct{}) {
	if n <= 0 { // 从n开始倒数
		return
	}
	ticker := time.NewTicker(1 * time.Second) // 创建一个周期性的定时器，每隔一秒
	for {
		countCh <- n // 把n放入管道
		<-ticker.C   // 等1秒
		n--          // n 减1
		if n <= 0 {
			ticker.Stop()          // 停止定时器
			finishCh <- struct{}{} // 成功结束
			break                  // 退出for循环
		}
	}
}

// 中止
func abort(ch chan struct{}) {
	buffer := make([]byte, 1)
	os.Stdin.Read(buffer)
	ch <- struct{}{}
}
func main() {
	countCh := make(chan int)
	finishCh := make(chan struct{})
	go countDown(countCh, 10, finishCh) // 开一个子协程，去往countCh和finishCh
	abourtCh := make(chan struct{})
	go abort(abourtCh) // 开一个子协程,往abortCh 里面放数据

	//LOOP:
	//	for {
	//		select {
	//		case n := <-countCh:
	//			fmt.Println(n)
	//		case <-finishCh:
	//			fmt.Println("finish")
	//			break LOOP // 退出For循环，在使用for select时，单独一个break不能退
	//		case <-abourtCh:
	//			fmt.Println("Abort")
	//			break LOOP // 退出for循环
	//
	//		}
	//	}
	// 使用return 可以直接退出 for + select 循环
	for {
		select {
		case n := <-countCh:
			fmt.Println(n)
		case <-finishCh:
			fmt.Println("finish")
			return
		case <-abourtCh:
			fmt.Println("Abort")
			return
		}
	}
}
