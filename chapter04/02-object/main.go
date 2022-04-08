package main

import (
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
)

func main() {
	persons := []Person{
		{
			name:   "TF",
			sex:    "男",
			tall:   1.72,
			weight: 65,
			age:    28,
		},
	}
	tf := Person{
		name:   "TF",
		sex:    "男",
		tall:   1.72,
		weight: 65,
		age:    28,
	}
	fmt.Println(tf)
	for _, item := range persons {
		bmi, err := gobmi.BMI(item.weight, item.tall)
		fmt.Println(bmi, err)
	}

	a := new(Person)
	// 相当于 a := &Person{}
	fmt.Println(a)
}

type human struct {
}
type Person struct {
	name   string
	sex    string
	tall   float64
	weight float64
	age    int
}
