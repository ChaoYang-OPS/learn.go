package main

import (
	"fmt"
	"sync"
)

func main() {
	for i:=0; i < 10;i++{
		//countDictGorouteSafe()
		countDictForGetUnlLock()  // fatal error: all goroutines are asleep - deadlock!
	}
}

func countDict() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	wg := sync.WaitGroup{}
	wg.Add(5000)
	for p := 0; p < 5000; p++ {
		go func() {
			defer wg.Done()
			//fmt.Println("正在统计第", p, "页")
			totalCount += 100 // totalCount = totalCount + 100   注意， 这里有重复覆盖的问题
		}()
	}
	wg.Wait()
	fmt.Println("预计有",100*5000,"字")
	fmt.Println("总共有:", totalCount, "字")
}

func countDictGorouteSafe() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	totalCountLock := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(5000)
	for p := 0; p < 5000; p++ {
		go func() {
			defer wg.Done()
			//fmt.Println("正在统计第", p, "页")
			totalCountLock.Lock()
			// 也可以用defer totalCountLock.Unlock()解锁，推荐
			totalCount += 100 // totalCount = totalCount + 100   注意， 这里有重复覆盖的问题
			totalCountLock.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("预计有",100*5000,"字")
	fmt.Println("总共有:", totalCount, "字")
}

func countDictForGetUnlLock() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	totalCountLock := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(5)
	for p := 0; p < 5; p++ {
		go func() {
			defer wg.Done()
			//fmt.Println("正在统计第", p, "页")
			totalCountLock.Lock()
			totalCount += 100 // totalCount = totalCount + 100   注意， 这里有重复覆盖的问题
			//totalCountLock.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("预计有",100*5000,"字")
	fmt.Println("总共有:", totalCount, "字")
}