package main

import "fmt"

func main() {
	var names [3]string
	var sexs [3]int
	var talls [3]float64
	var weights [3]float64
	var ages [3]int
	var bmis [3]float64
	var faterates [3]float64
	for i := 0; i < 3; i++ {
		var name string
		fmt.Print("请输入姓名:")
		fmt.Scanln(&name)
		names[i] = name
		var sex int
		fmt.Print("请输入性别:(男/女->1/0):")
		fmt.Scanln(&sex)
		sexs[i] = sex
		var tall float64
		fmt.Print("请输入身高(米):")
		fmt.Scanln(&tall)
		talls[i] = tall
		var weight float64
		fmt.Print("请输入体重:")
		fmt.Scanln(&weight)
		weights[i] = weight
		var age int
		fmt.Print("请输入年龄:")
		fmt.Scanln(&age)
		ages[i] = age
		var bmi float64 = weight / (tall * tall)
		bmis[i] = bmi
		if sex == 1 {
			sex = 1
		} else {
			sex = 0
		}
		var faterate float64 = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sex)) / 100
		fmt.Println(faterate)
		faterates[i] = faterate

		if sex == 1 {
			if age >= 18 && age <= 39 {
				if faterate > 0.0 && faterate < 0.1 {
					fmt.Println("偏塑")
				} else if faterate > 0.1 && faterate < 0.16 {
					fmt.Println("标准")
				} else if faterate > 0.16 && faterate < 0.20 {
					fmt.Println("偏重")
				} else if faterate > 0.25 && faterate < 0.30 {
					fmt.Println("肥胖")
				} else {
					fmt.Println("严重肥胖")
				}
			} else if age >= 40 && age <= 59 {
				if faterate > 0.0 && faterate < 0.11 {
					fmt.Println("偏塑")
				} else if faterate > 0.11 && faterate < 0.17 {
					fmt.Println("标准")
				} else if faterate > 0.17 && faterate < 0.22 {
					fmt.Println("偏重")
				} else if faterate > 0.22 && faterate < 0.27 {
					fmt.Println("肥胖")
				} else {
					fmt.Println("严重肥胖")
				}
			} else {
				if faterate > 0.0 && faterate < 0.13 {
					fmt.Println("偏塑")
				} else if faterate > 0.13 && faterate < 0.19 {
					fmt.Println("标准")
				} else if faterate > 0.19 && faterate < 0.24 {
					fmt.Println("偏重")
				} else if faterate > 0.24 && faterate < 0.29 {
					fmt.Println("肥胖")
				} else {
					fmt.Println("严重肥胖")
				}
			}
		} else {
			if age >= 18 && age <= 39 {
				if faterate > 0.0 && faterate < 0.20 {
					fmt.Println("偏塑")
				} else if faterate > 0.21 && faterate < 0.27 {
					fmt.Println("标准")
				} else if faterate > 0.27 && faterate < 0.34 {
					fmt.Println("偏重")
				} else if faterate > 0.35 && faterate < 0.39 {
					fmt.Println("肥胖")
				} else {
					fmt.Println("严重肥胖")
				}
			} else if age >= 40 && age <= 59 {
				if faterate > 0.0 && faterate <= 0.21 {
					fmt.Println("偏塑")
				} else if faterate > 0.21 && faterate <= 0.28 {
					fmt.Println("标准")
				} else if faterate > 0.28 && faterate <= 0.35 {
					fmt.Println("偏重")
				} else if faterate > 0.35 && faterate <= 0.40 {
					fmt.Println("肥胖")
				} else {
					fmt.Println("严重肥胖")
				}
			} else {
				if faterate > 0.0 && faterate <= 0.22 {
					fmt.Println("偏塑")
				} else if faterate > 0.22 && faterate <= 0.29 {
					fmt.Println("标准")
				} else if faterate > 0.29 && faterate <= 0.36 {
					fmt.Println("偏重")
				} else if faterate > 0.36 && faterate <= 0.41 {
					fmt.Println("肥胖")
				} else {
					fmt.Println("严重肥胖")
				}
			}
		}
	}
	var total float64
	for i := 0; i < 3; i++ {
		fmt.Println(names[i], ages[i], talls[i], weights[i])
		total += faterates[i]
	}
	fmt.Println("avge:", total/3)
}
