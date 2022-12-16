package apps

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn.go/testNe/rest-demo-host/apps/host"
)

// IOC 容器层: 管理所有的服务的实例

// 1. HostService的实例必须注册过来, HostService才会有具体的实例, 服务启动时注册
// 2. HTTP 暴露模块, 依赖Ioc里面的HostService

var (
	// 40 Service, 怎么办? 写40个定义,
	// 使用Interface{} + 断言进行抽象
	HostService host.Service

	// 维护当前所有的服务
	implApps = map[string]ImplService{}
	ginApps  = map[string]GinService{}
)

type ImplService interface {
	Config()
	Name() string
}

func RegistryImpl(svc ImplService) {
	if _, ok := implApps[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s has registried", svc.Name()))
	}
	implApps[svc.Name()] = svc
	if v, ok := svc.(host.Service); ok {
		HostService = v
	}
}

// 如果指定了具体类型, 就导致没增加一种类型, 多一个Get方法
// func GetHostImpl(name string) host.Service

// Get 一个Impl服务的实例：implApps
// 返回一个对象, 任何类型都可以, 使用时, 由使用方进行断言

func GetImpl(name string) interface{} {
	for k, v := range implApps {
		if k == name {
			return v
		}
	}
	return nil
}

func RegistryGin(svc GinService) {
	// 服务实例注册到svcs map当中
	if _, ok := ginApps[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s has registried", svc.Name()))
	}
	ginApps[svc.Name()] = svc
}

type GinService interface {
	Registry(r gin.IRouter)
	Config()
	Name() string
}

// 用户初始化 注册到Ioc容器里面的所有服务

func InitImpl() {
	for _, v := range implApps {
		v.Config()
	}
}

// 已经加载完成的Gin App由哪些

func LoadedGinApps() (names []string) {
	for k := range ginApps {
		names = append(names, k)
	}
	return
}

// 用户初始化 注册到Ioc容器里面的所有服务
func InitGin(r gin.IRouter) {
	// 先初始化好所有对象
	for _, v := range ginApps {
		v.Config()
	}
	// 完成Http Handler的注册
	for _, v := range ginApps {
		v.Registry(r)
	}
}
