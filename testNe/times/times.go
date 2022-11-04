package main

import (
	"fmt"
	"time"
)

func timeOfBaseUse() {
	// 时间差用Sub
	// 当前时间跟t0的时间差 用Since
	// time.Add()
	// time.After 判断当前时间晚于传入时间
	// 定时执行
	// time.NewTimer
	// time.After
	// 周期执行
	// time.NewTicker(1*time.Second)
}

func ticker() {
	// 周期性执行某个任务
	tk := time.NewTicker(1 * time.Second)
	for i := 0; i < 3; i++ {
		<-tk.C
		fmt.Println(time.Now().Unix())
	}
	defer tk.Stop()
}
func timer() {
	tm := time.NewTimer(time.Second)
	<-tm.C
	fmt.Println(time.Now().Unix())
	tm.Stop()
}
func format() {
	var a int = 6
	fmt.Printf("%b\n", a)
	type A struct {
		Name string
		Age  int
	}
	b := A{}
	b.Name = "TF"
	b.Age = 18
	fmt.Printf("%T\n", b)
	fmt.Printf("%v\n", b)
	fmt.Printf("%+v\n", b)
	fmt.Printf("%#v\n", b)
}
func main() {
	//now := time.Now()
	//fmt.Println(now)
	//da, _ := now.MarshalJSON()
	//fmt.Println(string(da))
	layout := "2006-01-02 15:04:05"
	now := time.Now()
	nowToString := now.Format(layout)
	fmt.Println(nowToString) // 2022-11-02 21:51:52
	local, _ := time.LoadLocation("Asia/Shanghai")
	// stringToTime  Recommend add LoadLocation
	t, err := time.ParseInLocation(layout, nowToString, local)
	if err != nil {
		return
	}
	fmt.Println(t.UnixMicro(), t.UnixNano()) // 1667397624000000
	nowTmp := time.Now()
	fmt.Println(nowTmp.Hour(), nowTmp.Minute(), nowTmp.Second())
	fmt.Println(nowTmp.Day(), now.YearDay())
	fmt.Println(nowTmp.Unix())
	fmt.Println(nowTmp.UnixMicro())
	fmt.Println(nowTmp.UnixMilli())
	fmt.Println(nowTmp.UnixNano())
	fmt.Println(nowTmp.Weekday())
	fmt.Println(nowTmp.Weekday().String())

	time.Sleep(3 * time.Second)
	t2 := time.Now()
	durations := t2.Sub(now)
	durations2 := time.Since(now)
	fmt.Println(durations.Seconds())
	fmt.Println(durations2.Seconds())
	t3 := now.Add(durations)
	fmt.Println(t3.Unix())
	ticker()
	timer()
	format()
}
