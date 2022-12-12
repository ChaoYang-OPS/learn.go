package onces

import (
	"fmt"
	"sync"
	"testing"
)

//   sync.Once用来确保某个动作至多执行一次
//  普遍用于初始化资源，单例模式
type MyBiz struct {
	once sync.Once
}

func (m *MyBiz) Init() {
	m.once.Do(func() {
		fmt.Println("I am is once")
	})
}

type singleon struct {
}

var instance *singleon

var instanceOnce sync.Once

func GetSingleInstance() *singleon {
	instanceOnce.Do(func() {
		instance = &singleon{}
		fmt.Println("....")
	})
	return instance
}
func TestGetSingleInstance(t *testing.T) {
	for i := 0; i < 10; i++ {
		GetSingleInstance()
	}
}
