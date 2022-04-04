package calc

// shift + command + - 函数合并

func CalcFatRate(bmi float64, age int, sex string) (fatRate float64) {
	sexWeight := 0
	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}
	fatRate = (1.2*bmi + getAgeWeight(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	return
}

func getAgeWeight(age int) (ageWeight float64) {
	ageWeight = 0.23
	if age >= 30 && age <= 40 {
		ageWeight = 0.22
	}
	return
}
