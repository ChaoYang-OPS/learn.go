package main

import "fmt"

type shipLegend struct {
}

func (*shipLegend) OpenTheDoorOfRefrigerator(refrigerator Refrigerator) error {
	fmt.Println("用ship做OpenTheDoorOfRefrigerator")
	return nil
}
func (*shipLegend) PutElephantIntoRefrigerator(elephant Elephant, refrigerator Refrigerator) error {
	fmt.Println("用ship做PutElephantIntoRefrigerator")
	return nil
}
func (*shipLegend) CloseTheDoorOfRefrigerator(refrigerator Refrigerator) error {
	fmt.Println("用ship做CloseTheDoorOfRefrigerator")
	return nil
}
