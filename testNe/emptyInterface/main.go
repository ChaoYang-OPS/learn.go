package main

import "fmt"

//  空接口没有定义任何方法，因此任意类型都实现了空接口
type C interface{}

type A struct {
	Name string
}

type B struct {
	Price float32
	A     // 形式上相当B继承了A
}

func main() {
	var c C // C是个空接口
	a := A{}
	b := B{}
	_ = b.Name
	c = a
	c = b
	fmt.Println(c)

	var i int = 5
	var bo bool = true
	c = i
	c = bo
	fmt.Println(c)

	arr := make([]interface{}, 0, 100)
	arr = append(arr, 9)
	arr = append(arr, 9.0)
	arr = append(arr, "tes")
	arr = append(arr, true)
	fmt.Println(arr)
	mp := make(map[interface{}]interface{}, 100)
	mp[3] = "2"
	mp["r4"] = "45"
	mp[A{}] = B{}
	fmt.Println(mp)
}
