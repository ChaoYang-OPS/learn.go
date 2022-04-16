package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		//countDictGorouteSafe()
		//countDictForGetUnlLock()  // fatal error: all goroutines are asleep - deadlock!
		countDictGorouteSafeRW()
		// 1:7:6
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
	fmt.Println("预计有", 100*5000, "字")
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
	fmt.Println("预计有", 100*5000, "字")
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
	fmt.Println("预计有", 100*5000, "字")
	fmt.Println("总共有:", totalCount, "字")
}

func countDictGorouteSafeRW() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	totalCountLock := sync.RWMutex{}
	wg := sync.WaitGroup{}
	workCount := 5
	wg.Add(workCount)
	doneCh := make(chan struct{})
	for p := 0; p < workCount; p++ {
		go func(p int) {
			result := map[int64]struct{}{}
			for { // 读锁可以多个GoRouting同时拿到
				fmt.Println(p, "读锁开始时间", time.Now())
				totalCountLock.RLock()
				fmt.Println(p, "读锁拿到锁时间", time.Now())
				fmt.Println(totalCount)
				result[totalCount] = struct{}{}
				totalCountLock.RUnlock()
				needBreak := false
				select {
				case <-doneCh:
					needBreak = true
				default:

				}
				if needBreak {
					//break
					return
				}
			}
			fmt.Println("result.....", result)
		}(p)
	}

	for p := 0; p < workCount; p++ {
		go func(p int) {
			defer wg.Done()

			fmt.Println(p, "写锁开始时间", time.Now())
			//fmt.Println("正在统计第", p, "页")
			totalCountLock.Lock()
			fmt.Println(p, "写锁拿到锁时间", time.Now())
			// 也可以用defer totalCountLock.Unlock()解锁，推荐
			totalCount += 100 // totalCount = totalCount + 100   注意， 这里有重复覆盖的问题
			totalCountLock.Unlock()
		}(p)
	}
	wg.Wait()
	close(doneCh)
	fmt.Println("预计有", 100*5000, "字")
	fmt.Println("总共有:", totalCount, "字")
}
