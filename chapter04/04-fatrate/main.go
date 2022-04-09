package main

import "fmt"

func main() {
	person := getFakePersonInfo()
	c := Calc{}
	c.BMI(person)
	c.FatRate(person)
	fmt.Println(person)

	frSvc := &fatRateService{
		s: GetFatRateSuggestion(),
	}
	frSvc.s = GetFatRateSuggestion()
	fakePerson := getFakePersonInfo()
	fmt.Println(frSvc.GiveSuggestionToPerson(fakePerson))
}

func getPersonInfoFromInput() (*Person) {
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
	var sex string
	fmt.Print("性别(男/女)：")
	fmt.Scanln(&sex)
	return &Person{
		name: name,
		sex: sex,
		tall: tall,
		weight: weight,
		age: age,
	}
}

func getFakePersonInfo () *Person {
	return &Person{
		name: "TF",
		sex: "男",
		tall: 1.7,
		weight: 70,
		age: 35,
	}
}