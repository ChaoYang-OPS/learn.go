package main

type PutElephantIntoRefrigerator interface {
	OpenTheDoorOfRefrigerator(refrigerator Refrigerator) error
	PutElephantIntoRefrigerator(elephant Elephant, refrigerator Refrigerator) error
	CloseTheDoorOfRefrigerator(refrigerator Refrigerator) error
}

type Refrigerator struct {
	Size string
}

type Elephant struct {
	Name string
}
