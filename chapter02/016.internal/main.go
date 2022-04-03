package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	//arr := []int{}
	startTime := time.Now()
	defer func() {
		finishTime := time.Now()
		fmt.Println(startTime, finishTime)
		//  查看运行了多长时间
		fmt.Println(finishTime.Sub(startTime))
	}()
	defer func() {
		fmt.Println("1")
	}()
	defer func() {
		fmt.Println("2")
	}()
	defer func() {
		fmt.Println("3")
	}()
	arr2 := testPanic()
	i := new(int)
	j := 333
	k := &j
	fmt.Println(reflect.TypeOf(i)) // *int
	fmt.Println(reflect.TypeOf(k)) // *int
	arr3 := make([]string, 3, 4)
	//for i := 0; i<len(arr2);i++{
	//	arr3[i] = arr2[i]
	//}
	copy(arr3, arr2)
}

func testPanic() []string {
	arr2 := make([]string, 0, 4)
	fmt.Println(len(arr2), cap(arr2)) // 0 4
	fmt.Println("default", arr2[0])
	fmt.Println("default", arr2[1])
	fmt.Println("default", arr2[2])
	return arr2
}
