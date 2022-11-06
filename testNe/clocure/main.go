package main

import (
	"fmt"
	"time"
)

func main() {

	arr := []int{1, 2, 3, 4}
	for _, ele := range arr {
		go func(value int) {
			fmt.Printf("%d\n", value)
		}(ele)
	}
	time.Sleep(1 * time.Second)
}
