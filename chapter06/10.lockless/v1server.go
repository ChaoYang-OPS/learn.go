package main

import (
	"fmt"
	"sync"
	"time"
)

type WebServerV1 struct {
	config     Config
	configLock sync.RWMutex
}

func (ws *WebServerV1) Visit() string {
	ws.configLock.Lock()
	defer ws.configLock.Unlock()
	return ws.config.Content
}

func (ws *WebServerV1) Reload() {
	ws.configLock.Lock()
	defer ws.configLock.Unlock()
	ws.config.Content = fmt.Sprintf("%v", time.Now())
}

func (ws *WebServerV1) ReloadWorker() {
	for {
		time.Sleep(10 * time.Millisecond)
		ws.Reload()
	}
}
