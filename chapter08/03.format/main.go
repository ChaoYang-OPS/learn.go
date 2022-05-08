package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"learn.go/pkg/apis/api"
	"os"
)

func main() {
	filePath := "./test-write.txt"
	persionalInformation := api.PersonalInformation{
		Name:   "TF",
		Sex:    "男",
		Tall:   1.68,
		Weight: 62,
		Age:    21,
	}
	fmt.Printf("%+v\n", persionalInformation)
	data, err := json.Marshal(persionalInformation)
	if err != nil {
		return
	}
	fmt.Println("Marshal的结果是原始", data)
	fmt.Println("Marshal的结果是string", string(data))
	writeFile(filePath, data)
	readFile(filePath)
}

func readFile(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取文件失败:", err)
		return
	}
	fmt.Println("读取出来的内容是:", string(data))
	//infos := strings.Split(string(data), ",")
	personalInformation := api.PersonalInformation{}
	fmt.Println("开始计算体脂信息:", string(data))
	json.Unmarshal(data, &personalInformation) // todo handle error
	fmt.Println("开始计算体脂信息:", personalInformation)
}
func writeFile(filePath string, data []byte) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("无法打开文件")
		os.Exit(1)
	}
	defer file.Close()
	file.Write(data)
}
