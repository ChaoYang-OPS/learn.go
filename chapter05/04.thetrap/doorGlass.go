package main

import "fmt"

// 关键点
var _ Door = &GlassDoor{}

// 很重要
type GlassDoor struct{}

func (d *GlassDoor) Unlock() {
	fmt.Println("GlassDoor Unlock")
}

func (d *GlassDoor) Lock() {
	fmt.Println("GlassDoor Lock")
}

func (*GlassDoor) Open() {
	fmt.Println("GlassDoor Open")
}

func (*GlassDoor) Close() {
	fmt.Println("GlassDoor Close")
}
