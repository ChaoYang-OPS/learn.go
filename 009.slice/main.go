package main

import "fmt"

func main() {
	a := []int{} // a是切片
	fmt.Println(a)
	fmt.Println("追加元素到a中，a是切片")
	a = append(a, 3333)
	fmt.Println(a)
	b := [0]int{} // b是数组
	fmt.Println(b)
	xqInfo := []string{"TF", "ll", "online"}
	fmt.Println(xqInfo)
	for i, v := range xqInfo {
		fmt.Println(i, v)
	}
	fmt.Println(xqInfo[0])

	fmt.Println("==============================================")
	fmt.Println("删除切片中的元素")
	a = []int{111, 222, 333, 444, 555}
	fmt.Println("删除前：", a) //删除前： [111 222 333 444 555]
	a = append(a[:2], a[3:]...)
	// a = append(a[:2],a[2:4]...)
	fmt.Println("删除后：", a) // 删除后： [111 222 444 555]
	a = append(a, a[:]...)
	fmt.Println("double", a) // double [111 222 444 555 111 222 444 555]
	// 备份
	backup := append([]int{}, a[1:]...)
	a = append(a[:1], 999)
	a = append(a, backup...)
	fmt.Println(a) // 	fmt.Println("double", a) // double [111 222 444 555 111 222 444 555]
}
