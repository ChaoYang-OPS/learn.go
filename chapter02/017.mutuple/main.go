package main

import "fmt"

func main() {
	bmis := []float64{1.2, 3.222, 4.3443}
	avgBmi := calculateAvg(bmis...)
	rs := calculateAvgOnSlice(bmis)
	fmt.Println(avgBmi)
	fmt.Println(rs)
	getScoreOfStudent("Tom")
}

func calculateAvg(bmis ...float64) (avgBmi float64) {
	var total float64 = 0
	for _, item := range bmis {
		total += item
	}
	avgBmi = total / float64(len(bmis))
	return
}

func calculateAvgOnSlice(bmis []float64) float64 {
	var total float64 = 0
	for _, item := range bmis {
		total += item
	}
	return total / float64(len(bmis))
}

func getScoreOfStudent(stdName string) (chinese int, math int, english int, physic int, nature int) {
	return 0, 0, 0, 0, 0
}
