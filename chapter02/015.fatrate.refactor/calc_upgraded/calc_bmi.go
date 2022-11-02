package calc

func CalcBMI(tall float64, weight float64) (bmi float64) {
	if tall <= 0 {
		panic("身高不能小于等于0，或者是负数")
	}
	// todo  验证体重的合法性
	return weight / (tall * tall)
}
