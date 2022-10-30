package main

import (
	"fmt"
	"sync"
)

type User struct {
	Name string
	Age  int
}

func NewUser() *User {
	return &User{
		Name: "",
		Age:  -1,
	}
}

type PeoPle struct {
	Name string
	Age  int
}

var (
	user     *User
	userOnce sync.Once

	people     *PeoPle
	peopleOnce sync.Once
)

type Config struct {
	Os     string
	Mem    int
	DBConn string
}

func GetUserInstance() *User {
	userOnce.Do(func() {
		if user == nil {
			user = NewUser()
		}
	})
	return user
}

func GetPeopleInstance() *PeoPle {
	peopleOnce.Do(func() {
		if people == nil {
			people = new(PeoPle)
		}
	})
	return people
}
func main() {
	u1 := GetUserInstance()
	u2 := GetUserInstance()
	fmt.Printf(" u1 %p\n", u1) // 0xc00000c030
	fmt.Printf(" u2 %p\n", u2) // 0xc00000c030
	p := GetPeopleInstance()
	p2 := GetPeopleInstance()
	fmt.Printf("%p\n", p)
	fmt.Printf("%p\n", p2)
}
