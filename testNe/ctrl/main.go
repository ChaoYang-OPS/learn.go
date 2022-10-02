package main

import (
	"fmt"
	"time"
)

func goo() {
	a, b, c, d := 1, 2, 3, 4
	if a > 0 {
		if b > 0 {
			if c > 0 {
				if d > 0 {
					fmt.Println("OK")
				}
			}
		}
	}
	// 规避多层嵌套
	if (a <= 0) || (3 > 6 && d > 7) {
		return
	}
	if b <= 0 {
		return
	}
	if c <= 0 {
		return
	}

	if d <= 0 {
		return
	}
	fmt.Println("OK")
	m := make(map[int]string, 10)
	m[1] = "abc"
	fmt.Printf("[%s]\n", m[2])
	if value, exist := m[2]; exist {
		fmt.Printf("[%s]\n", value)
	} else {
		value = "23"
		exist = false
	}
}
func ifNest() {
	now := time.Now()
	weekday := now.Weekday()
	day := now.Day()
	hour := now.Hour()
	a, b, c := now.Date()
	fmt.Println(weekday, hour, day)
	fmt.Println(a, b, c)
	if weekday.String() != "Saturday" && weekday.String() != "Sunnday" {
		if hour <= 9 && hour >= 7 {
			fmt.Println("公交车道不让走")
		} else {
			fmt.Println("可以走")
		}
	} else {
		fmt.Println("可以走.")
	}
}
func main() {
	ifNest()
	fmt.Println("Hello World")
}
