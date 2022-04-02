package main

import "fmt"

func main() {
	names := []string{"TF", "luck", "Tom"}
	fr := []float64{28, 17, 18}
	names = append(names, "SRE")
	fr = append(fr, 19)
	for i, name := range names {
		if name == "luck" {
			fmt.Println(fr[i])
		}
	}
	fmt.Println("定义map")
	var m1 map[string]int = nil
	// 必须实例化才可以写入
	//m1["a"] = 1  // panic on nil map
	delete(m1, "a")
	fmt.Println("没有实例化，直接取数", m1["a"])
	m2 := map[string]int{}
	m3 := map[string]int{"Tom": 66, "TF": 98, "Sun": 55}
	fmt.Println(m1, m2, m3)

	fmt.Println("Tom的分数：", m3["Tom"]) //66
	fmt.Println("TF的分数：", m3["TF"])   // 98
	fmt.Println("强哥的分数：", m3["强哥"])   // 0
	qGscore, ok := m3["强哥"]
	fmt.Println(qGscore, ">>>>>>", ok)
	m3["强哥"] = 77
	fmt.Println("强哥的分数：", m3["强哥"]) // 77
	qGscore, ok = m3["强哥"]
	fmt.Println(qGscore, ">>>>>>", ok)
	for name, score := range m3 {
		fmt.Printf("%s\t%d\n", name, score)
	}
}
