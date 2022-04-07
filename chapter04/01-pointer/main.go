package main

import "fmt"

func main() {
	a, b := 1, 2
	add(&a, &b)
	fmt.Println(a) // 3
	c := &a        // c 类型是*int ,c 指向a的盒子， *c就可以拿到a里面的东西 3
	d := &c        // d的类型是**int ，d指向c的盒子， d本身是指针，它存的东西也是指针
	fmt.Println(*c)
	fmt.Println("d=", d, "*d=", *d, "**d=", **d) // d= 0xc00000e030 *d= 0xc000018058 **d= 3
	m := map[string]string{}
	mp1 := &m // mp1的类型就是 *map[string]string
	fmt.Println(mp1)
	put(m)
	fmt.Println("*mp1=", *mp1) //*mp1= map[a:AAAA]
	f1 := add                  // f1 = func(int, int)
	f1(&a, &b)
	fmt.Println("f1, add= ", a) // f1, add=  5
	f1p1 := &f1                 // f1p1 = *func(int, int)
	(*f1p1)(&a, &b)
	fmt.Println("f1p1, add= ", a) // f1p1, add=  7

	var nothing *int
	//*nothing = 3  // 注意，这里是没有指向任何东西的int指针  panic: runtime error: invalid memory address or nil pointer dereference
	fmt.Println(nothing) // 	nil
	{
		var nothingMap map[string]string
		//nothingMap["a"] = "BBB" // panic: assignment to entry in nil map
		fmt.Println(nothingMap)
	}
}

func put(m map[string]string) {
	m["a"] = "AAAA"
}
func add(a, b *int) {
	*a = *a + *b
}
