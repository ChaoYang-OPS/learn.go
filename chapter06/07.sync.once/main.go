package main

import "sync"

type Rank struct {
	standard []string
}

var globalRank = &Rank{}

var globalRankInitialized bool = false
var globalRankInitializedLock sync.Mutex

func init() {
	globalRank.standard = []string{"Asia"}
}

func initGlobalRankStandard(standard []string) {
	globalRankInitializedLock.Lock()
	defer globalRankInitializedLock.Unlock()

	if globalRankInitialized {
		return
	}
	globalRank.standard = standard
	globalRankInitialized = true
}

func main() {
	standard :=[]string{"Asia"}
	for i := 0; i < 10;i++ {
		go func() {
			initGlobalRankStandard(standard)
		}()
	}
}