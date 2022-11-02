package prime

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRunPrime(t *testing.T) {
	startTime := time.Now()
	result := []int{}
	for num := 2; num <= 200000; num++ {
		if isPrine(num) {
			result = append(result, num)
		}
	}
	finsishTime := time.Now()
	fmt.Println(len(result))
	fmt.Println("共耗时", finsishTime.Sub(startTime))
}

func TestRunPrime2(t *testing.T) {
	startTime := time.Now()
	result := []int{}
	//wg := sync.WaitGroup{}
	//wg.Add(2)
	go func() {
		//defer wg.Done()
		fmt.Println("第一个work开始计算", time.Now())
		result = append(result, collectPrine(2, 100000)...)
		fmt.Println("第一个work完成计算", time.Now())
	}()
	go func() {
		//defer wg.Done()
		fmt.Println("第二个work开始计算", time.Now())
		result = append(result, collectPrine(100001, 200000)...)
		fmt.Println("第二个work开完成计算", time.Now())
	}()
	time.Sleep(15 * time.Second)
	//wg.Wait()
	finsishTime := time.Now()
	fmt.Println(len(result))
	fmt.Println("共耗时", finsishTime.Sub(startTime))
}

func TestRunPrime3(t *testing.T) {
	startTime := time.Now()
	result := []int{}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("第一个work开始计算", time.Now())
		result = append(result, collectPrine(2, 100000)...)
		fmt.Println("第一个work完成计算", time.Now())
	}()
	go func() {
		defer wg.Done()
		fmt.Println("第二个work开始计算", time.Now())
		result = append(result, collectPrine(100001, 200000)...)
		fmt.Println("第二个work开完成计算", time.Now())
	}()
	wg.Wait()
	finishTime := time.Now()
	fmt.Println(len(result))
	fmt.Println("共耗时", finishTime.Sub(startTime))
}

func collectPrine(start int, end int) (result []int) {
	for num := start; num <= end; num++ {
		if isPrine(num) {
			result = append(result, num)
		}
	}
	return
}

func isPrine(num int) (isPrine bool) {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrine = false
			return
		}
	}
	isPrine = true
	return
}

func TestHelloGoroutine(t *testing.T) {
	go fmt.Println("hello goroutine")
}
func TestHelloGoroutine2(t *testing.T) {
	go fmt.Println("hello goroutine")
	time.Sleep(1 * time.Second)
}

func TestForLoop(t *testing.T) {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 100; i < 120; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

}

//func TestForverGoroutine(t *testing.T) {
//	go func() {
//	for	{
//			time.Sleep(1 * time.Second)
//			go func() {
//				fmt.Println("启动新的GoRoutine@",time.Now())
//				time.Sleep(1*time.Hour)
//			}()
//		}
//	}()
//	for {
//		fmt.Println(runtime.NumGoroutine())
//		time.Sleep(2*time.Second)
//	}
//}
