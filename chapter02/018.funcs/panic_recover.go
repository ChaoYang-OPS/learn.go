package main

import (
	"fmt"
	"runtime/debug"
)

func panicAndRecover() {
	defer coverPanicUpgraded() // 这样是能抓住严重错误的

	defer func() {   // 这个是抓不住严重错误的
		coverPanicUpgraded()
	}()

	var nameScore map[string]int = nil
	nameScore["TF"] = 100
}

func coverPanic() { // 未能抓住panic
	func() {
		if r := recover(); r != nil {
			fmt.Println("系统出现重大故障", r)
		}
	}()
}

func coverPanicUpgraded() {
	if r := recover(); r != nil {
		fmt.Println("系统出现重大故障", r)
		debug.PrintStack()
	}
}
