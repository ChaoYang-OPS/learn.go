package main

import (
	"fmt"
)

func main() {
	var left, right int = 1, 2
	//var op string = "+"
	c := Calculator{
		left:  left,
		right: right,
		//op: op,
	}
	fmt.Printf("&c @main %p\n", &c)
	fmt.Println(c.Add())               // 调用add函数，c的result= 3
	fmt.Println("c.result=", c.result) // c.result= 3

	//newC := NewCalculater{&Calculator{}}
	newC := NewCalculater{}
	newC.left = 100
	newC.right = 200
	fmt.Println(newC.Add())
	// 初始化结构体，实例化map
	mc := MyCommand{commandOptions: map[string]string{},mainCommnad: new(string)}
	mc.commandOptions["aa"] = "AAA"  // panic: assignment to entry in nil map
	*mc.mainCommnad = "test"
	fmt.Println(*mc.mainCommnad)  // test

	fmt.Println(mc.ToCmdStr())
}

type MyCommand struct {
	mainCommnad *string
	commandOptions map[string]string
}

func (my MyCommand) ToCmdStr() string {
	out := ""
	for k,v := range my.commandOptions {
		out = out + fmt.Sprintf("---%s-=%s\n",k,v)
	}
	return out
}