package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	filePath := "./test-write.txt"
	writeFile(filePath)
	readFile(filePath)
}

func readFile(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取文件失败:", err)
		return
	}
	infos := strings.Split(string(data), ",")
	fmt.Println("开始计算体脂信息:", string(data))
	fmt.Println(infos)
}
func writeFile(filePath string) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("无法打开文件")
		os.Exit(1)
	}
	defer file.Close()
	file.Write([]byte("TF，男，1.70,..."))
}
