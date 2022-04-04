package main

import (
	"fmt"
	v1 "learn.go/chapter02/015.fatrate.refactor/calc"
	v2 "learn.go/chapter02/015.fatrate.refactor/calc_upgraded"
)

func init() {
	fmt.Println("我是init函数1")
}
func main() {

	// 选中shift+(
	for {
		mainFatRateBody()
		if cont := whetherContinue(); !cont {
			return
		}
	}
}
func init() {
	fmt.Println("我是init函数2")
}
func mainFatRateBody() {
	weight, tall, age, sexWeight, _ := getMaterialsFromInput()
	//fatRate := calc.CalcFatRate()
	fatRate := calcFatRate(weight, tall, age, sexWeight)
	// 回调函数
	getHealthinessSuggestions(age, fatRate, getHealthinessSuggestionForMale)
	getHealthinessSuggestions(age, fatRate, getHealthinessSuggestionForFMale)
	// check 特殊变量
	check := getHealthinessSuggestionForFMale
	check(age, fatRate)
}

//  getSuggestion 工具
func getHealthinessSuggestions(age int, fatRate float64, getSuggestion func(age int, fatRate float64)) {
	getSuggestion(age, fatRate)
}

func generateCheckHealthinessForMale() func(age int, fatRate float64) {
	// 定制
	return func(age int, fatRate float64) {

	}
}

func getHealthinessSuggestionForMale(age int, fatRate float64) {
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
}
func getHealthinessSuggestionForFMale(age int, fatRate float64) {
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
}

func calcFatRate(weight float64, tall float64, age int, sexWeight int) float64 {
	// 计算体脂率
	//var bmi float64 = weight / (tall * tall)
	bmi := v2.CalcBMI(tall, weight)
	_ = v1.CalcBMI(tall, weight)
	var fatRate float64 = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	fmt.Println("体脂率是:", fatRate)
	return fatRate
}

func getMaterialsFromInput() (float64, float64, int, int, string) {
	var name string
	fmt.Print("姓名：")
	fmt.Scanln(&name)
	var weight float64
	fmt.Print("体重（KG）：")
	fmt.Scanln(&weight)
	var tall float64
	fmt.Print("身高（米）：")
	fmt.Scanln(&tall)
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
	return weight, tall, age, sexWeight, sex
}
func whetherContinue() bool {
	var whetherConting string
	fmt.Print("是否录入下一个(y/n)?")
	fmt.Scanln(&whetherConting)
	if whetherConting != "y" {
		return false
	}
	return true
}
