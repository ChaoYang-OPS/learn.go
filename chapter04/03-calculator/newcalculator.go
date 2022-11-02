package main

type NewCalculater struct {
	//*Calculator 如果是嵌套的指针，则在实例化NewCalculater的时候，必须实例化嵌套的实体，否则会空指针
	Calculator
}
