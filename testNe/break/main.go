package main

import "fmt"

func breakFor() {
	arr := []int{1, 2, 3, 4, 5}
	for i, item := range arr {
		fmt.Println(i, item)
		if i > 2 {
			break
		}
		fmt.Printf("next %d\n", i+1)
	}
}

func continueFor() {
	arr := []int{1, 2, 3, 4, 5}
	for i, item := range arr {
		fmt.Println(i, item)
		if i > 2 {
			continue
		}
		fmt.Printf("next %d\n", i+1)
	}
}
func complexBreakContiunue() {
	const SIZE = 5
	arr := [SIZE][SIZE]int{}
	for i := 0; i < SIZE; i++ {
		fmt.Printf("第%d行\n", i)
		if i%2 == 1 {
			for j := 0; j < SIZE; j++ {
				fmt.Printf("第%d列\n", j)
				if arr[i][j]%2 == 0 {
					continue
				}
				fmt.Printf("将要第%d列\n", j+1)
			}
			break
		}
	}
}
func main() {
	//breakFor()
	continueFor()
	complexBreakContiunue()
}
