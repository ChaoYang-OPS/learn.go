package main

import (
	"fmt"
	"math/rand"
)

func NewRandFirst() {
	source := rand.NewSource(1)
	render := rand.New(source)
	resultFirst := make([]int, 0, 10)
	resultSecond := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", render.Intn(100))
		resultFirst = append(resultFirst, render.Intn(100))
	}
	// 81 87 47 59 81 18 25 40 56 0
	fmt.Println()
	source.Seed(1)
	render2 := rand.New(source)
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", render2.Intn(100))
		resultSecond = append(resultSecond, render2.Intn(100))
	}
	fmt.Println()
	fmt.Println(resultFirst)
	fmt.Println(resultSecond)
	// 81 87 47 59 81 18 25 40 56 0

	arr := rand.Perm(10)
	fmt.Println(arr) // [9 4 2 6 8 0 3 1 7 5]
	rand.Shuffle(10, func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	// [8 6 9 5 3 7 2 4 1 0]
	fmt.Println(arr)
}
func main() {
	NewRandFirst()
}
