package main

import "fmt"

var conter int

func calcSum(nums ...int) (sum int) {
	for _, item := range nums {
		sum += item
	}
	conter++
	return
}

func showUsedTimes() {
	fmt.Println("used:", conter)
}

func genImprovementFunc() func(percentage float64) {
	base := 1000.0
	return func(percentage float64) {
		base = base * (1 + percentage)
		fmt.Println(base)
	}
}

func closureMain() {
	imp := genImprovementFunc()
	imp(0.1)
	imp(0.1)
	imp(0.1)
	imp(0.1)
}
