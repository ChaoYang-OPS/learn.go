package main

import "fmt"

func subSlice() {
	arr := make([]int, 3, 5)
	fmt.Println(arr)
	crr := arr[0:2]
	crr[1] = 8
	fmt.Println(arr[1])
	//fmt.Println(arr)  // [0 8 0]
	crr = append(crr, 9)
	fmt.Println(arr[2])
	crr = append(crr, 9)
	crr = append(crr, 9)
	// 内存分离
	crr = append(crr, 9)
	fmt.Println(arr)        //[0 8 9]
	fmt.Println(crr)        // [0 8 9 9 9 9]
	fmt.Printf("%p\n", arr) // 0xc0000aa060
	fmt.Printf("%p\n", crr) // 0xc0000b6000
}
func main() {
	subSlice()
	var ch chan int
	fmt.Println("ch is nil", ch == nil)
	fmt.Println(len(ch))
	ch = make(chan int, 10)
	fmt.Println(len(ch))

	fmt.Println(len(ch))
	fmt.Println(cap(ch)) // 10
	for i := 0; i < 10; i++ {
		ch <- 4
	}
	// 取出元素,直接抛弃
	<-ch
	// 放入元素
	ch <- 3 // 死锁
	fmt.Println(len(ch))
	fmt.Println(cap(ch))
	// 遍历之前必须先关闭channel,如果没关系管道，只能通过 for i:=0;
	//close(ch)
	for i := 0; i < len(ch); i++ {
		// 这种方式不要求先关闭管道
		elem := <-ch
		fmt.Println(elem)
	}
	fmt.Println("++++++++++++++++++++++++++++")
	//for elem := range ch {
	// //必须需要关闭管道
	//	// 并取走元素,等价于上面for循环
	//	fmt.Println(elem)
	//}
	fmt.Println(len(ch))
	fmt.Println(cap(ch))
}
