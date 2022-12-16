package protocol

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"learn.go/testNe/rest-demo-host/apps"
	"learn.go/testNe/rest-demo-host/configs"
	"log"
	"net/http"
	"time"
)

type HttpService struct {
	server *http.Server
	l      *log.Logger
	r      gin.IRouter
}

func NewHttpService() *HttpService {
	// new gin router实例, 并没有加载Handler
	r := gin.Default()

	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1M
		Addr:              configs.C().App.HttpAddr(),
		Handler:           r,
	}
	return &HttpService{
		server: server,
		l:      log.Default(),
		r:      r,
	}
}

func (s *HttpService) Start() error {
	// 加载Handler, 把所有的模块的Handler注册给了Gin Router
	apps.InitGin(s.r)
	// 已加载App的日志信息
	apps := apps.LoadedGinApps()
	//s.l.Printf("loaded gin apps :%v\n", apps)
	println(apps)
	// 该操作时阻塞的, 简单端口，等待请求
	// 如果服务的正常关闭

	if err := s.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			s.l.Println("service stoped success")
			return nil
		}
		return fmt.Errorf("start service error, %s", err.Error())
	}
	return nil
}
func (s *HttpService) Stop() {
	s.l.Println("start graceful shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		s.l.Printf("shut down http service error, %s\n", err)
	}
}
