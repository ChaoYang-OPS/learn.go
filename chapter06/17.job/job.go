package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Minute)
	defer cancel()
	successFlag := make(chan bool, 1)

	go account(ctx)
	go distributeService(ctx)
	go configure(ctx)
	go verifyService(ctx, successFlag)
	select {
	case <-ctx.Done():
		fmt.Println("超时，没有完成")
	case v := <-successFlag:
		if v {
			fmt.Println("任务完成，成功结束")
		} else {
			fmt.Println("任务失败，需要重新重试还是下线服务")
		}
	}
}

func account(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	doneCh := make(chan string, 2)
	go accountRegister(ctx, doneCh)
	go accountGrantAuth(ctx, doneCh)
	successCount := 0
	for v := range doneCh {
		successCount++
		fmt.Println("job:", v, "done")
		if successCount == 2 {
			close(doneCh)
		}
	}
	fmt.Println("账号处理完成")
}

func accountRegister(ctx context.Context, doneCh chan string) {
	fmt.Println("注册账号")
	defer fmt.Println("注册账号完成")
	for { //  call .......
		select {
		case <-ctx.Done():
			fmt.Println("context结束，不在注册")
			return
		default:
		}
		// .... 调用xxx接口
		doneCh <- "accountRegister"
		fmt.Println("accountRegister 成功")
		break
	}
}
func accountGrantAuth(ctx context.Context, doneCh chan string) {
	fmt.Println("授权账号")
	defer fmt.Println("授权账号完成")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context结束，不在授权")
			return
		default:
		}
		// .... 调用xxx接口
		doneCh <- "accountGrantAuth"
		fmt.Println("accountGrantAuth 成功")
		break
	}
}

func distributeService(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 7*time.Minute)
	defer cancel()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go distributeLB(ctx, &wg)
	go distributeInstance(ctx, &wg)
	wg.Wait()
	fmt.Println("distributeService done")
}

func distributeInstance(ctx context.Context, w *sync.WaitGroup) {
	defer w.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("上下文结束，要删除已经创建的实例")
			return
		default:
		}
		fmt.Println("部署实例")
		break
	}
}

func distributeLB(ctx context.Context, w *sync.WaitGroup) {
	defer w.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("上下文结束，要删除已经创建的LB")
			return
		default:
		}
		fmt.Println("部署LB")
		break
	}
}

func configure(_ context.Context) {
	fmt.Println("注入新服务账号")
}

func verifyService(ctx context.Context, flag chan bool) {
	ctx, cancle := context.WithCancel(ctx)
	defer cancle()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go verifyFunction(ctx, &wg)
	wg.Wait()
	fmt.Println("验证服务完成")
	flag <- true
}

func verifyFunction(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("上下文结束，不再验证")
			return
		default:
		}
		fmt.Println("开始验证服务")
		time.Sleep(1 * time.Second) // 用来替换验证部分的环节，比如服务调用，服务模拟等
		if i <= 1 {
			fmt.Println("服务尚未完成，重试")
		}
		fmt.Println("服务验证完成")
		break
	}
}
