package main

import (
	"fmt"
	"sync"
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

var facStore = &dbFactoryStore{}

type Conn struct{}

type dbFactoryStore struct {
	store map[string]DBFactory
}
type DBFactory interface {
	GetConnection() *Conn
}

func initMySqlFac(connStr string) DBFactory {
	return &MySqlDBFactory{}
}

type MySqlDBFactory struct {
	once sync.Once
}

func (MySqlDBFactory) GetConnection() *Conn {
	once.Do(func() {
		initMySqlFac("")
	})
	// todo
	return nil
}

var counter int = 0
var counterOnce sync.Once

type School struct {
	classroomLocation map[string]string
}

func main() {
	//standard := []string{"Asia"}
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		initGlobalRankStandard(standard)
	//	}()
	//}
	//time.Sleep(1 * time.Second)
	for i := 0; i < 10; i++ {
		fmt.Println("第x次", i)
		counterOnce.Do(func() {
			fmt.Println("初始化")
			counter++
		})
	}
	fmt.Println("最终结果:", counter)
}
