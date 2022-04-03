package main

import "fmt"

func main() {
	hello()
	helloToSomeone("TF")
	helloToSomeone("Tom")
	msg := constructHelloMessage("The")
	fmt.Println(msg)
}

func hello() {
	fmt.Println("您好，golang")
}

func helloToSomeone(name string) {
	fmt.Println("您好：", name)
}

func constructHelloMessage(name string) string {
	//fmt.Println("您好：", name)
	return "您好" + name
}

func calcBMI(tall, weight float64) float64 {
	return tall / (weight * weight)
}
