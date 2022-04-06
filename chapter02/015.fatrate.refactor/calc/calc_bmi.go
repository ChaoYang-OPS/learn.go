package calc

import "fmt"

func CalcBMI(tall float64, weight float64) (bmi float64, err error) {
	if tall <= 0 {
		//panic("身高不能小于等于0，或者是负数")
		return 0, fmt.Errorf("身高不能小于等于0，或者是负数")
	}
	bmi = weight / (tall * tall)
	// todo  验证体重的合法性
	return
}
