package main

import "fmt"

func main() {
	fmt.Println("1rd version")
	// 难以长期维护
	xgInfo := []string{"Luckily", "男", "online"}
	xlInfo := []string{"TF", "男", "online"}
	xsInfo := []string{"TFS", "男", "online"}
	fmt.Println(xgInfo)
	fmt.Println(xlInfo)
	fmt.Println(xsInfo)
	// ....
	// 难点，数组长度管理
	fmt.Println("2rd version")
	newPersonInfos := [3][3]string{ // 难点
		[3]string{"Luckily", "男", "online"},
		[3]string{"TF", "男", "online"},
		[3]string{"TFS", "男", "online"},
	}
	for _, val := range newPersonInfos {
		fmt.Println(val)
	}
	// ... 支持动态添加
	fmt.Println("3rd version")
	newPersonInfos2 := [...][3]string{ // 难点
		[3]string{"Luckily", "男", "online"},
		[3]string{"TF", "男", "online"},
		[3]string{"TFS", "男", "online"},
	}
	for _, val := range newPersonInfos2 {
		fmt.Println(val)
	}
	fmt.Println("------------------------------------------------")
	// 用降维方式输出
	for d1, d1val := range newPersonInfos2 {
		for d2, d2value := range d1val {
			fmt.Println(d1, d1val, d2, "val:", d2value)
		}
	}
}
