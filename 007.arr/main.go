package main

import "fmt"

func main() {
	var age int = 35
	fmt.Println(age)
	var ages [5]int = [5]int{35, 32, 33, 37, 59}
	fmt.Println(ages)
	var ages2 = [5]int{35, 32, 33, 37, 59}
	fmt.Println(ages2)
	ages3 := [5]int{35, 32, 31, 33, 37}
	fmt.Println(ages3)
	//ages3 = [6]int{35,32,31,33,37,99}  错误： 数组长度不能变
	ages2 = ages3
	var ages4 [3]int
	fmt.Println(ages4) // [0 0 0]
	ages4[0] = 100
	ages4[1] = 1111
	ages4[2] = 2222
	fmt.Println(ages4) // 	[100 1111 2222]

	for i := 0; i < len(ages4); i++ {
		fmt.Println(ages4[i])
	}
	for i, val := range ages4 {
		fmt.Println(ages4[i], "====>", i, "---->", val)
	}
}
