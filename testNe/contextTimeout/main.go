package main

import (
	"context"
	"fmt"
	"time"
)

const (
	WorkUseTime = 500 * time.Microsecond
	Timeout     = 100 * time.Microsecond
)

// 模拟一个耗时比较长的任务
func LongTimeWork() {
	time.Sleep(WorkUseTime)
	return
}

// 模拟一个接口处理函数
func Handle1() {
	deadline := make(chan struct{}, 1)
	workDone := make(chan struct{}, 1)
	go func() { // 要把控制超时的函数放到一个协程里
		LongTimeWork()
		workDone <- struct{}{}
	}()

	go func() { // 要把控制超时的函数放到一个协程里
		time.Sleep(Timeout)
		deadline <- struct{}{}
	}()
	select {
	case <-workDone:
		fmt.Println("LongTimeWork return")
	case <-deadline:
		fmt.Println("LongTimeWork timeout")
	}
}

func Handle2() {
	workDone := make(chan struct{}, 1)
	go func() {
		LongTimeWork()
		workDone <- struct{}{}
	}()
	select {
	// 下面的case只执行最早到来的那一个
	case <-workDone:
		fmt.Println("LongTimeWork return")
	case <-time.After(Timeout):
		fmt.Println("LongTimeWork timeout by Handle2 ")
	}
}

func Handle3() {
	ctx, cancel := context.WithCancel(context.Background())
	workDone := make(chan struct{}, 1)
	go func() {
		LongTimeWork()
		workDone <- struct{}{}
	}()
	go func() {
		// 100毫秒后调用cancel，关闭ctx.Done()
		time.Sleep(Timeout)
		cancel()
	}()
	select {
	case <-workDone:
		fmt.Println("LongTimeWork return")
	case <-ctx.Done(): // ctx.Done()是一个管道,context超时或者调用了	cancel()
		fmt.Println("LongTimeWork timeout by Handle3 ")
	}
}
func main() {
	Handle1()
	Handle2()
	Handle3()
}
