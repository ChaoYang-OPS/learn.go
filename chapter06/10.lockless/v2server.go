package main

import (
	"fmt"
	"time"
)

type WebServerV2 struct {
	config *Config
}

func (ws *WebServerV2) Visit() string {
	return ws.config.Content
}

func (ws *WebServerV2) Reload() {
	ws.config = &Config{
		Content: fmt.Sprintf("%v", time.Now()),
	}
}

func (ws *WebServerV2) ReloadWorker() {
	for {
		ws.Reload()
	}
}
