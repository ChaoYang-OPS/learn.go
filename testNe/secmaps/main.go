package main

import (
	"fmt"
	"sync"
	"time"
)

type Student struct {
	age  int
	name string
}

var (
	lst     []int = make([]int, 5)
	arr     [5]int
	student Student
	//mp      map[int]int = make(map[int]int)
	mp sync.Map
)

func rwShareMemory1() {
	for i := 1; i < len(lst); i += 2 {
		lst[i] = 555
	}
	student.age = 18
	// 并发写入MAP 会报错
	//for i := 0; i < 10; i++ {
	//	mp[i] = i
	//}
	for i := 0; i < 100; i++ {
		// 写入MAP
		mp.Store(i, i)
	}
}

func rwShareMemory2() {
	for i := 0; i < len(lst); i += 2 {
		lst[i] = 888
	}
	student.name = "TFS"
	//for i := 0; i < 10; i++ {
	//	mp[i] = i * i
	//}
	//for i := 0; i < 100; i++ {
	//	// 写入MAP
	//	mp.Store(i, i)
	//}
	for i := 0; i < 100; i++ {
		// 读取
		mp.Load(i)
	}
}

func main() {
	go rwShareMemory1()
	go rwShareMemory2()
	time.Sleep(1 * time.Second)
	fmt.Println(lst)
	fmt.Println(student) // {18 TFS}
	// fatal error: concurrent map writes
	//fmt.Println(mp)
	fmt.Println(mp.Load(1)) // 1 true
	fmt.Println(mp.Load(3)) // 3 true
}
