package main

import "fmt"

type User struct {
	Name string
	Age  int
	Sex  byte
}

func NewUser() *User {
	return &User{
		Name: "",
		Age:  0,
		Sex:  0,
	}
}
func main() {
	up := new(User)
	fmt.Println("Hello world", NewUser(), up)
}
