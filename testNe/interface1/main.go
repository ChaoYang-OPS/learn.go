package main

import "fmt"

type Transporter interface {
	move(int, int) (int, error)
	whilst() error
}

// 一个structs可以实现多个接口
// 接口的本质

// 接口值有两部分组成，一个指向该接口的具体类型的指针和另外一个指向该具体真实数据的指针
type Car struct {
	price int
	power float32
}
type Ship struct {
	price int
	power float32
}

func (car Car) whilst() error {
	fmt.Println("whils")
	return nil
}
func (car Car) move(a, b int) (int, error) {
	return a + b, nil
}
func (car Car) walk() {
	fmt.Println("walk....")
}

func (car Ship) whilst() error {
	fmt.Println("whils", car.price)
	return nil
}
func (car Ship) move(a, b int) (int, error) {
	return a + b, nil
}
func (car Ship) walk() {
	fmt.Println("walk....")
}
func foo(a Transporter) {
	a.whilst()
}

type A struct {
	Name string
}
type B struct {
	Price float32
	A     // 形式上相当B继承了A
}

type C interface {
	foo()
}

type D interface {
	good()
	C
}

func main() {
	var ifs Transporter
	var c Car
	var ccc Car
	var s Ship
	ifs = c
	ifs.whilst() // 实际是在执行c的whilst()
	ifs.move(1, 2)
	c.move(1, 2)
	c.walk()
	foo(c)
	foo(ccc) // 值实现的方法，指针也同样实现了
	foo(&s)
	sp := &Car{}
	sp.whilst()
	sp.walk()
}
