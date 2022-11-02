package main

import (
	"fmt"
	"sort"
)

type Button struct {
	Floor int
}

type Elevator struct {
	buttons   Buttons
	postition int
}

type Buttons []*Button

func (b Buttons) Len() int {
	return len(b)
}

func (b Buttons) Less(i, j int) bool {
	return b[i].Floor < b[j].Floor
}

func (b Buttons) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func main() {
	ev := &Elevator{
		postition: 2,
		buttons: Buttons{
			{Floor: 3},
			{Floor: 1},
			{Floor: 5},
			{Floor: 2},
			{Floor: 4},
			{},
		},
	}
	//sort.Sort(ev.buttons)
	sort.Sort(sort.Reverse(ev.buttons))
	fmt.Printf("%+v", ev.buttons)
	for _, item := range ev.buttons {
		fmt.Println(item.Floor)
	}
}

func sortButtons([]*Button) []*Button {
	return nil
}
