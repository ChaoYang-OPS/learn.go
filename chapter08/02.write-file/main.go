package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "./test-write.txt"
	//writeFile(filePath)
	writeFileWithAppend(filePath)
}

func writeFile(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("无法创建文件", filePath, "错误信息是:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.Write([]byte("this is first line---"))
	file.Write([]byte("this is first line\n"))
	file.WriteString("第二行内容\n")
	file.WriteAt([]byte("CHANGE"), 0)
	//file.Sync() 告诉操作系统直接把内容写到磁盘
}

func writeFileWithAppend(filePath string) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("无法打开文件")
		os.Exit(1)
	}
	defer file.Close()
	file.Write([]byte("this is first line---"))
	file.Write([]byte("this is first line\n"))
	file.WriteString("第二行内容\n")
	file.WriteAt([]byte("CHANGE"), 0)
	//file.Sync() 告诉操作系统直接把内容写到磁盘
}
