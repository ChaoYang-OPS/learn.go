package main

import "fmt"

func main() {
	var data string
	{
		var equipment IOInterface = &Soft{}
		data = equipment.Read()
		fmt.Println(data)
	}
	{
		var equipment IOInterface = &Mag{}
		data = equipment.Read()
		fmt.Println(data)
	}
	{
		var equipment IOInterface = &Paper{}
		data = equipment.Read()
		fmt.Println(data)
	}
	{
		var equipment IOInterface = &Sata{}
		data = equipment.Read()
		fmt.Println(data)
	}
}

type Sata struct {
}

func (Sata) Read() string {
	return "Sata"
}
func readFrom14Sort() string {
	return "lala"
}

func readFromPaper() string {
	return "........."
}

type Paper struct {
}

func (Paper) Read() string {
	return "paper"
}

type Soft struct {
}

type Mag struct {
}

func (Mag) Read() string {
	return "mag"
}
func (Soft) Read() string {
	return "软"
}

type IOInterface interface {
	Read() string
}
