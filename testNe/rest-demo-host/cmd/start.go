package cmd

import (
	"github.com/spf13/cobra"
	"learn.go/testNe/rest-demo-host/apps"
	_ "learn.go/testNe/rest-demo-host/apps/all"
	"learn.go/testNe/rest-demo-host/configs"
	"learn.go/testNe/rest-demo-host/protocol"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (

	// pusher service config option
	confType string
	confFile string
	confETCD string
)

// 程序的启动时 组装都在这里进行
// 1.
// StartCmd represents the base command when called without any subcommands

var StartCmd = &cobra.Command{
	Use: "start",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 加载程序配置
		if err := configs.LoadConfigFromEnv(); err != nil {
			return err
		}

		// 初始化日志
		// 加载我们Host Service的实体类
		// host service 的具体实现
		// service := impl.NewHostServiceImpl()

		// 注册HostService的实例到IOC
		// 采用: _ "gitee.com/go-course/restful-api-demo-g7/apps/host/impl" 完成注册
		// apps.HostService = impl.NewHostServiceImpl()

		// 如何执行HostService的config方法
		// 因为 apps.HostService是一个host.Service接口, 并没有 保存实例初始化(Config)的方法
		apps.InitImpl()

		svc := newManager()
		ch := make(chan os.Signal, 1)
		// channel是一种复合数据结构, 可以当初一个容器, 自定义的struct make(chan int, 1000), 8bytes * 1024  1Kb
		// 如果没close gc是不会回收的
		defer close(ch)
		// Go为了并发编程设计的(CSP), 依赖Channel作为数据通信的信道
		// 出现了一个思路模式的转变:
		//    单兵作战(只有一个Groutine) --> 团队作战(多个Groutine 采用Channel来通信)
		//    main { for range channel }  这个时候的channel仅仅想到于一个缓存, 必须选择带缓存区的channl
		//    signal.Notify 当中一个Goroutine, g1
		//    go svc.WaitStop(ch) 第二Goroutine, g2
		//    g1 -- ch1 --> g2
		//    g1 <-- ch2 -- g2
		//    g1 数据发送给ch1, g2 读取channle中的数据, chanel 只要生成好了就能用, 如果channle关闭
		//    设计channel 使用数据的发送方负责关闭, 相当于表示挂电话
		//    for range   由range帮忙处理了 chnanel 关闭后， read的中断处理
		//    for v,err := <-ch { if(err == io.EOF) break }

		signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT)
		go svc.WaitStop(ch)
		return svc.Start()
	},
}

func newManager() *manager {
	return &manager{
		http: protocol.NewHttpService(),
		l:    log.Default(),
	}
}

type manager struct {
	http *protocol.HttpService
	l    *log.Logger
}

func (m *manager) Start() error {
	return m.http.Start()
}

func (m *manager) WaitStop(ch <-chan os.Signal) {
	for v := range ch {
		switch v {
		default:
			m.l.Println("receive signal : %s", v)
			m.http.Stop()
		}
	}
}
func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f", "etc/demo.toml", "demo api 配置文件路径")
	RootCmd.AddCommand(StartCmd)
}
