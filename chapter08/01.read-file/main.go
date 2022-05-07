package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "./test.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("无法打开文件", filePath, "错误信息是:", err)
		os.Exit(1)
	}
	defer file.Close()
	b := make([]byte, 1024, 1024)
	n, err := file.Read(b)
	if err != nil {
		fmt.Println("无法读取文件:", err)
		os.Exit(2)
	}
	fmt.Println("读出的内容:", b)  // 如果不转换的string类型，输出的是一长串数字
	fmt.Println("读出的内容:", string(b))
	fmt.Println("n的大小:", n)
	fmt.Println("读出的内容：",string(b[:n]))  // 一定要给后续程序使用时， 截取到实际读取到的数据，而不是全部，否则后续程序无效读取也作为正常数据处理
}
