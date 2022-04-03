package main

import (
	"fmt"
	"strconv"
)

func main() {
	name := "Hello"
	{
		name := "123"
		val, err := strconv.Atoi(name) // 123 <nil>
		fmt.Println(val, err)
	}
	age := 30
	fmt.Printf("%p\n", &age) // 0xc000018058
	age, time := 32, "时间"
	fmt.Printf("%p\n", &age) // 0xc000018058
	fmt.Println(name, age, time)
	{
		age := 3
		fmt.Printf("%p\n", &age) // 0xc0000b2020
	}

}
