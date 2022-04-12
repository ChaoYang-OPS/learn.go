package main

import "fmt"

type manLegend struct {
}

func (*manLegend) OpenTheDoorOfRefrigerator(refrigerator Refrigerator) error {
	fmt.Println("用manLegend做OpenTheDoorOfRefrigerator")
	return nil
}
func (*manLegend) PutElephantIntoRefrigerator(elephant Elephant, refrigerator Refrigerator) error {
	fmt.Println("用manLegend做PutElephantIntoRefrigerator")
	return nil
}
func (*manLegend) CloseTheDoorOfRefrigerator(refrigerator Refrigerator) error {
	fmt.Println("用manLegend做CloseTheDoorOfRefrigerator")
	return nil
}
