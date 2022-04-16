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

func main() {
	standard := []string{"Asia"}
	for i := 0; i < 10; i++ {
		go func() {
			initGlobalRankStandard(standard)
		}()
	}
	time.Sleep(1 * time.Second)
}
