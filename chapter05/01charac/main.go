package main

import (
	"fmt"
	"reflect"
)

// 接口只能定义方法，不允许有任何变量
// 接口是引用类型

// 接口实例化时，接口不能直接使用变量，必须指向实现
type PutElephantIntoRefrigerator interface {
	OpenTheDoorOfRefrigerator(refrigerator Refrigerator) error
	PutElephantIntoRefrigerator(elephant Elephant, refrigerator Refrigerator) error
	CloseTheDoorOfRefrigerator(refrigerator Refrigerator) error
}

type Refrigerator struct {
	Size string
}

type Elephant struct {
	Name string
}

type PutElephantIntoRefrigeratorImpl struct {
}

func (legent PutElephantIntoRefrigeratorImpl) OpenTheDoorOfRefrigerator(refrigerator Refrigerator) error {
	// todo
	fmt.Println("打开冰箱门")
	return nil
}
func (legent PutElephantIntoRefrigeratorImpl) PutElephantIntoRefrigerator(elephant Elephant, refrigerator Refrigerator) error {
	// todo
	fmt.Println("装进去")
	return nil
}
func (legent PutElephantIntoRefrigeratorImpl) CloseTheDoorOfRefrigerator(refrigerator Refrigerator) error {
	// todo
	fmt.Println("关门")
	return nil
}

//
type TestTypeImplInterface func()

func (t TestTypeImplInterface) OpenTheDoorOfRefrigerator(refrigerator Refrigerator) error {
	return nil
}
func (t TestTypeImplInterface) PutElephantIntoRefrigerator(elephant Elephant, refrigerator Refrigerator) error {
	return nil
}

func (t TestTypeImplInterface) CloseTheDoorOfRefrigerator(refrigerator Refrigerator) error {
	return nil
}
func main() {
	var refrigeratorgerator Refrigerator
	var epephant Elephant
	fmt.Println(epephant)
	//var ip *int  // 默认是nil
	//fmt.Println(*ip)  invalid memory address or nil pointer dereference
	//var putER PutElephantIntoRefrigerator  // 默认是nil
	//putER.OpenTheDoorOfRefrigerator(refrigeratorgerator)  // invalid memory address or nil pointer dereference
	// 接口必须指向实现
	// 必须实现所有的方法，才叫做实现了接口，否则不算
	// 不限类型
	// ******注意： 如果某个对象的成员方法有的在对象上，有的在对象指针上，那么在给接口赋值时，必须使用指针 ****
	// 同一个对象可以实现多个接口
	// 接口之间可以组合
	var legend PutElephantIntoRefrigerator = PutElephantIntoRefrigeratorImpl{}
	legend.OpenTheDoorOfRefrigerator(refrigeratorgerator)
	legend.PutElephantIntoRefrigerator(epephant, refrigeratorgerator)
	legend.CloseTheDoorOfRefrigerator(refrigeratorgerator)

	var o Open = Refrigerator{}
	var c Close = Refrigerator{}
	fmt.Println(o, c)
	// 从范围小的向范围大转换会成功
	// 从范围大的向范围小的转会失败
	var b Box = Refrigerator{}
	fmt.Println(b)
	c = b
	//b = c

	var i interface{}
	i = 3
	fmt.Println(reflect.TypeOf(i), "value:", i)
	i = 3.2
	fmt.Println(reflect.TypeOf(i), "value:", i)
}

type Open interface {
	Open() error
}

type Close interface {
	Close() error
}

func (Refrigerator) Open() error {
	return nil
}

func (Refrigerator) Close() error {
	return nil
}

type Box interface {
	Open
	Close
}
