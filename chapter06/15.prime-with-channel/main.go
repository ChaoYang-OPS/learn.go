package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	maxNum := 2000000
	result := make(chan int,maxNum)
	wg := sync.WaitGroup{}
	workerNumber := 16
	wg.Add(workerNumber)
	baseNumCh := make(chan int,1024)
	for i:=0 ; i< workerNumber;i++{
		go func() {
			defer wg.Done()
			for oNum := range baseNumCh{
				if isPrine(oNum) {
					result<-oNum
				}
			}
		}()
	}
	for num :=2; num <= maxNum; num ++ {
		baseNumCh<-num
	}
	close(baseNumCh)
	wg.Wait()
	finsishTime := time.Now()
	fmt.Println(len(result))
	fmt.Println("共耗时",finsishTime.Sub(startTime))
}
func isPrine(num int) (isPrine bool) {
	for i := 2; i < num ; i++ {
		if num %i == 0 {
			isPrine = false
			return
		}
	}
	isPrine =true
	return
}