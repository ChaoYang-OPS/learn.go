package main

import (
	"fmt"
	"testing"
	"time"
)

func TestDefChannel(t *testing.T) {
	var sampleMap map[string]int = map[string]int{}
	var intCh chan int // nil  引用类型
	fmt.Println(sampleMap, intCh)
	intCh = make(chan int, 1)
	fmt.Println(intCh) // 0xc0000741e0
	fmt.Println("装入数据")
	intCh <- 3
	fmt.Println("取出数据")
	out := <-intCh
	fmt.Println("数据是", out)
}

func TestChanPutGet(t *testing.T) {
	intCh := make(chan int) // 这是一个不带size的channel，不带buffer
	workerCount := 10
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			intCh <- i
		}(i)
	}
	for o := 0; o < workerCount; o++ {
		go func(o int) {
			out := <-intCh
			fmt.Printf("出%d拿到了%d\n", o, out)
		}(o)
	}
	time.Sleep(1 * time.Second)
}

// 这是一个让 out部分等待一段时间再取数据，来观察i的行为
// 结果： 如果没有out，那么in的部分会等待，直到有out开始
// 结果： 带有buffer的会直接允许in，不影响out
func TestChanPutGet2_Owait_withBuffer(t *testing.T) {
	intCh := make(chan int, 10) // 这是一个不带size的channel，不带buffer
	workerCount := 10
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			fmt.Println(i, "开始工作", time.Now())
			intCh <- i
			fmt.Println(i, "结束工作", time.Now())
		}(i)
	}
	fmt.Println(time.Now())
	time.Sleep(2 * time.Second)
	fmt.Println(time.Now())
	for o := 0; o < workerCount; o++ {
		go func(o int) {
			out := <-intCh
			fmt.Printf("%s 出%d拿到了%d\n", time.Now(), o, out)
		}(o)
	}
	time.Sleep(1 * time.Second)
}

//
func TestChanPutGet2_Owait_withSmallBuffer(t *testing.T) {
	intCh := make(chan int, 2) // 这是一个不带size的channel，不带buffer
	workerCount := 10
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			fmt.Println(i, "开始工作", time.Now())
			intCh <- i
			fmt.Println(i, "结束工作", time.Now())
		}(i)
	}
	fmt.Println(time.Now())
	time.Sleep(2 * time.Second)
	fmt.Println(time.Now())
	for o := 0; o < workerCount; o++ {
		go func(o int) {
			out := <-intCh
			fmt.Printf("%s 出%d拿到了%d\n", time.Now(), o, out)
		}(o)
	}
	time.Sleep(1 * time.Second)
}

// 先启动out，in后启动

func TestChanPutGet2_OFirst_withBuffer(t *testing.T) {
	intCh := make(chan int, 10) // 这是一个不带size的channel，不带buffer
	workerCount := 10
	for o := 0; o < workerCount; o++ {
		go func(o int) {
			out := <-intCh
			fmt.Printf("%s 出%d拿到了%d\n", time.Now(), o, out)
		}(o)
	}
	fmt.Println(time.Now())
	time.Sleep(2 * time.Second)
	fmt.Println(time.Now())
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			fmt.Println(i, "开始工作", time.Now())
			intCh <- i
			fmt.Println(i, "结束工作", time.Now())
		}(i)
	}
	time.Sleep(1 * time.Second)
}

// range没有close的话，在取出所有数据后 panic DeadLock
func TestRangeChannel(t *testing.T) {
	intCh := make(chan int, 10)
	intCh <- 1
	intCh <- 2
	intCh <- 3
	intCh <- 4
	intCh <- 5
	for o := range intCh {
		fmt.Println(o)
	}
}

func TestRangeClosedChannel(t *testing.T) {
	intCh := make(chan int, 10)
	intCh <- 1
	intCh <- 2
	intCh <- 3
	intCh <- 4
	intCh <- 5
	close(intCh)
	{
		o1, ok := <-intCh
		fmt.Println(o1, "直接取数:", o1, ok)
	}
	for o := range intCh {
		fmt.Println("Range取数", o)
	}
	o1, ok := <-intCh
	fmt.Println(o1, ok) // 0 false
}

func TestPut2ClosedChannel(t *testing.T) {
	intCh := make(chan int, 10)
	intCh <- 1
	intCh <- 2
	intCh <- 3
	intCh <- 4
	intCh <- 5
	close(intCh)
	intCh <- 5 // panic: send on closed channel
}

func TestOutOnly(t *testing.T) {
	intCh := make(chan int, 10)
	<-intCh // fatal error: all goroutines are asleep - deadlock!
}

func TestSingleGoroutineApp(t *testing.T) {
	intCh := make(chan int)
	intCh <- 1 // fatal error: all goroutines are asleep - deadlock!
	<-intCh    // fatal error: all goroutines are asleep - deadlock!
}

// 当所有channel不ready的时候，select会等待，直到某个channel ready
func TestSelectChannel(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan string)
	fmt.Println("start", time.Now())
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "123456"
	}()
	select {
	case o1 := <-ch1:
		fmt.Println(time.Now(), "ch1 ready,go", o1)
	case o2 := <-ch2:
		fmt.Println(time.Now(), "ch2 ready go", o2)
	}
	fmt.Println("Done")
}

// default，直接运行
func TestSelectChannelWithDefault(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan string)
	fmt.Println("start", time.Now())
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "123456"
	}()
	select {
	case o1 := <-ch1:
		fmt.Println(time.Now(), "ch1 ready,go", o1)
	case o2 := <-ch2:
		fmt.Println(time.Now(), "ch2 ready go", o2)
	default:
		fmt.Println(time.Now(), "所有的channel都没准备好，直接走default")
	}
	fmt.Println("Done")
}

// case的优先级高于default，只要有case中的channel ready ，default就不走
func TestSelectChannelWithDefaultAndChannelReady(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan string)
	fmt.Println("start", time.Now())
	//go func() {
	//	time.Sleep(1*time.Second)
	//	ch1 <-1
	//}
	// 写入数据
	ch1 <- 1
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "123456"
	}()
	select {
	case o1 := <-ch1:
		fmt.Println(time.Now(), "ch1 ready,go", o1)
	case o2 := <-ch2:
		fmt.Println(time.Now(), "ch2 ready go", o2)
	default:
		fmt.Println(time.Now(), "所有的channel都没准备好，直接走default")
	}
	fmt.Println("Done")
}

// closed channel 在select的时候总是可用，
func TestSelectChannelWithDefaultAndCloseChannel(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan string)
	fmt.Println("start", time.Now())
	//go func() {
	//	time.Sleep(1 * time.Second)
	//	ch1 <- 1
	//}()
	fmt.Println("关闭ch1")
	close(ch1)
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "123456"
	}()
	select {
	case o1 := <-ch1:
		fmt.Println(time.Now(), "ch1 ready,go", o1)
	case o2 := <-ch2:
		fmt.Println(time.Now(), "ch2 ready go", o2)
	default:
		fmt.Println(time.Now(), "所有的channel都没准备好，直接走default")
	}
	fmt.Println("Done")
}

func TestMutipleSelect(t *testing.T) {
	// 1:24:11
	ch1 := make(chan int)
	for i := 0; i < 10; i++ {
		go func(i int) {
			select {
			case <-ch1:
				fmt.Println(time.Now(), i)
			}
		}(i)
	}
	fmt.Println("关闭channel", time.Now())
	close(ch1)
	time.Sleep(1 * time.Second)
}

func TestMultipleChannelReadySelect(t *testing.T) {
	ch1, ch2 := make(chan int), make(chan int)
	close(ch1)
	close(ch2)
	ch1Counter, ch2Counter := 0, 0
	for i := 0; i < 10000000; i++ {
		select {
		case <-ch1:
			ch1Counter++
		case <-ch2:
			ch2Counter++
		}
	}
	fmt.Println("ch1Counter", ch1Counter)
	fmt.Println("ch2Counter", ch2Counter)
}
