package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func GenSlice(n int) []int {
	source := rand.NewSource(0)
	renders := rand.New(source)
	arr := make([]int, 0, n)
	for i := 0; i < n; i++ {
		arr = append(arr, renders.Intn(128))
	}
	return arr
}

func CountUniq(arr []int) int {
	m := make(map[int]bool, len(arr))
	for _, item := range arr {
		m[item] = true
	}
	return len(m)
}

func Concat(arr []int) string {
	// 传入nil并不会报错
	var stringsB strings.Builder
	for _, elem := range arr {
		stringsB.WriteString(strconv.Itoa(elem))
		// Itoa is equivalent to FormatInt(int64(i), 10).
		// FormatInt比Itoa会高效点
		stringsB.WriteString(strconv.FormatInt(int64(elem), 10))
		stringsB.WriteString(" ")
	}
	return strings.Trim(stringsB.String(), " ")
}

func main() {
	arr := GenSlice(100)
	cnt := CountUniq(arr)
	fmt.Println("uniq elem ", cnt)
	r := Concat(arr)
	fmt.Printf("[%s]\n", r)
}
