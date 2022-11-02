package main

import (
	"fmt"
)

// 全局变量
var tall float64
var weight float64

func main() {
	panicAndRecover()
	// close call
	closureMain()
	calcSum(1, 2, 3, 4, 5)
	showUsedTimes()
	deferGuess()
	openFile()
	//guess(1, 100)
	//tall, weight = 1.70, 70.0
	fmt.Println(fib(10))
	fmt.Println("done calc,sleep....")
	//time.Sleep(10 * time.Second)
	calculatorAdd := func(a, b float64) float64 {
		return a + b
	}
	result := calculatorAdd(1, 3)
	fmt.Println(result)
	{
		personTall := 1.81
		personWeight := 90.0
		calculatorAdd(personTall, personWeight)
		// suggestions
	}
	{
		personTall := 1.81
		personWeight := 90.0
		calculatorAdd(personTall, personWeight)
	}
	// personTall,personWeight 的有效范围为{}内部，外部无效
	fmt.Println(tall, weight)
	sampleSubdomain()
	sampleSubdomain2()
	fmt.Println("全局变量赋值前")
	caclAdd()
	tall, weight = 1.70, 70.0
	fmt.Println("全局变量赋值后")
	caclAdd()
	// 重新定义重名的局部变量
	tall, weight := 1.70, 71.0
	caclAdd() //
	fmt.Println(tall, weight)
	tall, weight = 200, 78.0
	caclAdd()

}

func caclAdd() float64 {
	//tall, weight := "1.70", 70.0
	fmt.Println(tall + weight)
	return 0
}

func funcDef(nums ...int) (addResult int) {
	for _, item := range nums {
		addResult += item
	}
	return
}

func funcDef1(nums ...int) int {
	sum := 0
	for _, item := range nums {
		sum += item
	}
	return sum
}

func sampleSubdomain() {
	name := "Luckily"             // 声明变量 name 值是Luckily
	fmt.Println("name is ", name) //name is  Luckily
	{
		fmt.Println("name is ", name) // name is  Luckily
		name = "KR"                   // 重新赋值给了name，所以name的值是KR
		fmt.Println("name is ", name) // KR
	}
	fmt.Println(">>>>name is ", name) // KR
}

func sampleSubdomain2() {
	name := "Luckily"             // 声明变量 name 值是Luckily
	fmt.Println("name is ", name) //name is  Luckily
	{
		name = "Luckily--update"
		fmt.Println("name is ", name) // name is  Luckily
		name := "KR"                  //
		fmt.Println("name is ", name)
	}
	fmt.Println(">>>>name is >>> ", name) // Luckily
	if name == "Luckily" {
		a := 3
		fmt.Println(a)
	} else {
		a := 4
		fmt.Println(a)
	}
}

//func sampleSubdomainIf() {
//	if v := caclBMI(); v == 0 {
//		fmt.Println(v)
//	} else {
//		fmt.Println(v)
//	}
//	//fmt.Println(v) 无效，v的有效范围为if block
//}

func sampleSubdomainFor() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello golang,", i)
	}
	//fmt.Println(i)  i的有效范围为for block
}

func fib(n uint) uint {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
