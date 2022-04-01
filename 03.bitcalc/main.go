package main

import "fmt"

func main() {
	a,b := 100,31
	fmt.Println(a^b)  //123
	fmt.Println(b^a)  //123

	arr :=[]int{3,3,4,4,5,6,5,6,7}
	result := -1
	for _, item := range  arr {
		if result < 0 {
			result = item
		} else  {
			result = result ^ item
		}
	}
	fmt.Println(result)  // 7
}
