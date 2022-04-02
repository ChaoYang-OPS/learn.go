package main

import "fmt"

func main() {
	var line1 [2]float64
	var line2 [2]float64
	line1[0] = 10
	line1[1] = 20
	line2[0] = 10
	line2[1] = 20

	if line1[0]/line1[1] == line2[0]/line2[1] {
		fmt.Println("平行")
	} else {
		fmt.Println("不平行")
	}
}
