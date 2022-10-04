package main

import (
	"fmt"
	"time"
)

type User struct {
	id         int
	Score      float32
	name, addr string
	enrollment time.Time
}
type People struct {
	Name string
	Age  int
}

// UserMap 可以添加任意方法 自定义类型添加方法
type UserMap map[int]People

func (um UserMap) Get(id int) People {
	return um[id]
}
func (um UserMap) Say(id int) {
	delete(um, id)
}

func VerifyUserMap() {
	data := UserMap{0: People{Name: "name", Age: 12}, 1: People{
		Name: "test",
		Age:  13,
	}}
	fmt.Println(data.Get(0))
	data.Say(0)
	fmt.Println(data)
}

type Student struct {
	Name string
	string
	int
}

func (s *Student) GetName() {
	fmt.Println("----------------GetName-----------", s.Name)
}

func (s Student) ChangeNameNotSuccess() {
	s.Name = "data"
}
func (s *Student) ChangeNameReally() {
	s.Name = "data"
}

//NewStudent  构造函数
func NewStudent(name, city string, age int) *Student {
	//var s Student
	//s.Name = name
	//s.string = city
	//s.int = age
	//return &s
	//s := &Student{name, city, age}
	//return s
	s := &Student{
		Name:   name,
		string: city,
		int:    age,
	}
	return s
}

func updateSlice(arr []Student) {
	arr[0].Name = "789"
}

type Video struct {
	Length int
	Auther Student
}

// 成员函数
func (user User) hello(name string) string {
	return "hello" + name + "...." + user.name
}
func Hello(name string, user User) string {
	return "hello" + name + "...." + user.name
}
func (User) think(name string) string {
	return "hello" + name
}
func basicUse() {
	//var abc int
	//var cdf string
	var ds User
	fmt.Printf("%d , Score %g, name [%s], addr [%s] enrollment %v\n", ds.id,
		ds.Score, ds.name, ds.addr, ds.enrollment)
	ws := User{
		Score: 100, name: "zcy",
	}
	ws.Score = 93.5
	ws.enrollment = time.Now()
	a := ws.id + 24
	fmt.Printf("a=%d\n", a)
	fmt.Println(ws)
	fmt.Printf("%+v\n", ds)
	fmt.Println(ws.enrollment)
	//fmt.Printf("%#v\n", ds)
}
func basicMethod() {
	ws := User{
		Score: 10, name: "test",
	}
	fmt.Println(ws.hello("test"))
	fmt.Println(Hello("test", ws))
	fmt.Println(ws.think("test"))
}
func main() {
	basicUse()
	basicMethod()
	VerifyUserMap()
	var s Student
	s.Name = "zcy"
	s.string = "441"
	s.int = 24
	fmt.Println(s)
	nes := NewStudent("test", "xa", 18)
	nes.GetName()
	nes.ChangeNameReally()
	s.ChangeNameNotSuccess()
	fmt.Println(s)
	fmt.Println(nes)
	arr := []Student{s, s, s, s}
	updateSlice(arr)
	fmt.Println(arr[0].Name)
	v := Video{25, s}
	v.Auther.Name = "789"
	fmt.Println(s.Name)
}
