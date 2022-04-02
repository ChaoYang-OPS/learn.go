package main

import "fmt"

func main() {
	var name string
	fmt.Print("姓名：")
	fmt.Scanln(&name)
	var weight float64
	fmt.Print("体重（KG）：")
	fmt.Scanln(&weight)
	var tall float64
	fmt.Print("身高（米）：")
	fmt.Scanln(&tall)
	var bmi float64 = weight / (tall * tall)
	var age int
	fmt.Print("年龄：")
	fmt.Scanln(&age)
	var sexWeight int
	var sex string
	fmt.Print("性别(男/女)：")
	fmt.Scanln(&sex)
	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}
	// 选中shift+(
	var fatRate float64 = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	fmt.Println("体脂率是:", fatRate)
	if sex == "男" {
		// todo 编写男性的体脂率与体脂状态表
		if age >= 18 && age <= 39 {
			if fatRate <= 0.1 {
				fmt.Println("目前是偏胖素，要多吃多练，增强体质")
			} else if fatRate > 0.1 && fatRate <= 0.16 {
				fmt.Println("目前是：标准，太棒了，要保持")
			} else if fatRate > 0.16 && fatRate <= 0.21 {
				fmt.Println("目前是偏胖，吃晚饭多散步，消化消化")
			} else if fatRate > 0.21 && fatRate <= 0.26 {
				fmt.Println("目前是肥胖，少吃点，多运动")
			} else {
				fmt.Println("目前是非常肥胖，赶紧健身游泳跑步，既刻开始")
			}
		} else if age >= 40 && age <= 59 {
			//todo
		} else if age >= 68 {
			// todo
		} else {
			fmt.Println("我们不看未成年人体脂率，变法太大，无法评判")
		}
	} else {
		// todo 编写女性的体脂率与体脂状态表
	}

}
