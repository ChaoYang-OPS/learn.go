package main

import (
	"fmt"
	"math/rand"
	"time"
)

func basic() {
	arr := [...]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(arr); i++ {
		fmt.Println(i)
	}
	fmt.Println("====================")
	for i := 0; i < len(arr); {
		fmt.Printf("%d %d\n", i, arr[i])
		i++
	}
	fmt.Println("==========================")
	i := 0
	for i < len(arr) {
		fmt.Printf("%d %d\n", i, arr[i])
		i++
	}

	sum := 0
	for i < len(arr) {
		sum += arr[i]
		fmt.Printf("%d %d\n", i, arr[i])
		i++
	}
	for sum, i := 0, 0; i < len(arr) && sum < 100; i++ {
		sum += arr[i]
		fmt.Printf("sum=%d\n", sum)
	}
}

func forRange() {
	// 引用类型，slice map channel都可以通过for range 来遍历
	arr := []int{1, 2, 3, 4, 5}
	for i, elem := range arr {
		fmt.Printf(" %d %d\n", i, elem)
	}
	str := "我会唱ABC" // 1 个汉字占3个byte
	fmt.Println(len(str))
	for i, elem := range str {
		fmt.Printf("%d %c\n", i, elem)
	}
	brr := []byte(str)
	fmt.Println(len(brr))
	for i, elem := range brr {
		fmt.Printf("%d %d\n", i, elem)
	}
	m := map[int]int{1: 1, 2: 2, 3: 3}
	for k, v := range m {
		fmt.Printf("%d=%d\n", k, v)
	}
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		// 写入数据
		ch <- i
	}
	// 必须先close,然后再使用for ... range
	close(ch)

	for elem := range ch {
		fmt.Printf("++++++++++++++++%T %v\n", elem, elem)
	}
	// 读取过之后，再读取不会有结果
	for elem := range ch {
		fmt.Printf("------------%T %v\n", elem, elem)
	}
}

func deadLoop() {
	for {
		fmt.Println("zzzz")
		time.Sleep(1 * time.Second)
	}
}

func matrixMultiply(A, B [4][4]int) [4][4]int {
	C := [4][4]int{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sum := 0
			for k := 0; k < 4; k++ {
				sum += A[i][k] * B[k][j]
			}
			C[i][j] = sum
		}
	}
	return C
}
func genMatrix() [4][4]int {
	A := [4][4]int{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			A[i][j] = rand.Int()
		}
	}
	return A
}
func main() {
	basic()
	forRange()
	//deadLoop()
	A := genMatrix()
	B := genMatrix()
	C := matrixMultiply(A, B)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%d ", C[i][j])
		}
		fmt.Println()
	}
	//fmt.Println(A[2][3])
}
