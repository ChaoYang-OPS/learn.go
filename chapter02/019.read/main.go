package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"learn.go/chapter02/015.fatrate.refactor/calc"
)

func main() {

	// 录入
	var (
		name   string
		sex    string
		tall   float64
		weight float64
		age    int
	)
	cmd := &cobra.Command{
		Use:   "helathcheck",
		Short: "体脂计算器，并给出建议",
		Long:  "体脂计算器，基于BMI。。。。。",
		Run: func(cmd *cobra.Command, args []string) {
			// 计算
			fmt.Println("name:", name)
			fmt.Println("sex:", sex)
			fmt.Println("tall:", tall)
			fmt.Println("weight:", weight)
			fmt.Println("age:", age)
			// 评估结果
			bmi ,err := calc.CalcBMI(tall, weight)
			if err != nil {
				return
			}
			fatRage := calc.CalcFatRate(bmi, age, sex)
			fmt.Println("FatRage:", fatRage)
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "姓名")
	cmd.Flags().StringVar(&sex, "sex", "", "性别")
	cmd.Flags().Float64Var(&tall, "tall", 0, "身高")
	cmd.Flags().Float64Var(&weight, "weight", 0, "体重")
	cmd.Flags().IntVar(&age, "age", 0, "年龄")
	cmd.Execute()
}
