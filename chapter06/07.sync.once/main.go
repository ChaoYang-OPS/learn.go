package main

import (
	"fmt"
	"sync"
	"time"
)

type Rank struct {
	standard []string
}

var globalRank = &Rank{}
var once sync.Once = sync.Once{}

func initGlobalRankStandard(standard []string) {
	// 只做一次
	once.Do(func() {
		globalRank.standard = standard
		fmt.Println("settings.....")
	})
}

func main() {
	standard :=[]string{"Asia"}
	for i := 0; i < 10;i++ {
		go func() {
			initGlobalRankStandard(standard)
		}()
	}
	time.Sleep(1*time.Second)
}