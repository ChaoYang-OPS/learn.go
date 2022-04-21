package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//withCancel()
	//withTimeout()
	//withValue()
	withDeadline()
}


func withDeadline() {
	now := time.Now()
	newTime := now.Add(1*time.Second)
	ctx,_ := context.WithDeadline(context.TODO(),newTime)
	go tv(ctx)
	go mobile(ctx)
	go game(ctx)
	time.Sleep(3*time.Second)
}
func tv(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("关电视")
			return
		default:
			fmt.Println("看电视")
			time.Sleep(300*time.Millisecond)
		}
	}
}

func mobile(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("关手机")
			return
		default:
			fmt.Println("看手机")
			time.Sleep(300*time.Millisecond)
		}
	}
}
func game(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("关游戏机")
			return
		default:
			fmt.Println("玩游戏机")
			time.Sleep(300*time.Millisecond)
		}
	}
}
func withValue() {
	//ctx := context.TODO()
	//ctx = context.WithValue(ctx,"1","钱包")
	ctx := context.WithValue(context.TODO(),"1","钱包")
	fmt.Println("withValue",ctx.Value("1"))
	fmt.Println("withValue",ctx.Value("2"))
	fmt.Println("withValue",ctx.Value("3"))
	fmt.Println("withValue",ctx.Value("4"))
	goToPapa(ctx)
	time.Sleep(5*time.Second)
}

func goToPapa(ctx context.Context) {
	ctx = context.WithValue(ctx,"2","充电包")
	go func(ctx context.Context) {
		time.Sleep(1*time.Second)
		fmt.Println("goToPapa",ctx.Value("1"))
		fmt.Println("goToPapa",ctx.Value("2"))
		fmt.Println("goToPapa",ctx.Value("3"))
		fmt.Println("goToPapa",ctx.Value("4"))
	}(ctx)
	goTomama(ctx)
}

func goTomama(ctx context.Context) {
	ctx = context.WithValue(ctx,"3","小夹克")
	go func(ctx context.Context) {
		time.Sleep(1*time.Second)
		fmt.Println("goTomama",ctx.Value("1"))
		fmt.Println("goTomama",ctx.Value("2"))
		fmt.Println("goTomama",ctx.Value("3"))
		fmt.Println("goTomama",ctx.Value("4"))
	}(ctx)
	goToGrandman(ctx)
}

func goToGrandman(ctx context.Context) {
	ctx = context.WithValue(ctx,"4","苹果")
	go func(ctx context.Context) {
		time.Sleep(1*time.Second)
		fmt.Println("goToGrandman",ctx.Value("1"))
		fmt.Println("goToGrandman",ctx.Value("2"))
		fmt.Println("goToGrandman",ctx.Value("3"))
		fmt.Println("goToGrandman",ctx.Value("4"))
	}(ctx)
	goToParty(ctx)
}

func goToParty(ctx context.Context) {
	go func(ctx context.Context) {
		time.Sleep(1*time.Second)
		fmt.Println("goToParty",ctx.Value("1"))
		fmt.Println("goToParty",ctx.Value("2"))
		fmt.Println("goToParty",ctx.Value("3"))
		fmt.Println("goToParty",ctx.Value("4"))
	}(ctx)
}
func withTimeout() {
	ctx ,_ := context.WithTimeout(context.TODO(),1*time.Second)
	fmt.Println("开始部署望远镜,发送信号")
	go distibuteMainFramp(ctx)
	go distibuteMainBody(ctx)
	go distibuteMainCover(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("任务超时，没有完成")
	}
	time.Sleep(20*time.Second)  // 等待20秒后收到任务取消消息
}

func distibuteMainFramp(ctx context.Context) {
	time.Sleep(10*time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("任务取消 distibuteMainFramp")
		return
	}
	fmt.Println("部署主框架")
}

func distibuteMainBody(ctx context.Context) {
	time.Sleep(10*time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("任务取消 distibuteMainBody")
		return
	}
	fmt.Println("部署distibuteMainBody")
}

func distibuteMainCover(ctx context.Context) {
	time.Sleep(10*time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("任务取消 distibuteMainCover")
		return
	}
	fmt.Println("部署distibuteMainCover")
}



func withCancel() {
	ctx := context.TODO()
	ctx, cancel := context.WithCancel(ctx)
	fmt.Println("做蛋挞，要买材料")
	go buyFlour(ctx)
	go buyFlour(ctx)
	go buyEgg(ctx)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("没电了，取消购买所有食材")
	// 当调用cancel后，所有由此上下文衍生的contxt都会被cancel
	cancel()
	time.Sleep(1 * time.Second)
}

func buyFlour(ctx context.Context) {
	fmt.Println("去买面")
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done(): // todo 介绍一下struct{}
		fmt.Println("收到消息，不买面了")
		return
	default:
	}
	fmt.Println("买面")
}

func buyOil(ctx context.Context) {
	fmt.Println("去买油")
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done(): // todo 介绍一下struct{}
		fmt.Println("收到消息，不买油了")
		return
	default:
	}
	fmt.Println("买油")
}

func buyEgg(ctx1 context.Context) {
	// defer cancel() 无论是否调用衍生出来的ctx的cancel，Done返回的channel都会关闭
	ctx, _ := context.WithCancel(ctx1)
	fmt.Println("去买蛋")
	//time.Sleep(1*time.Second)
	select {
	case <-ctx.Done(): // todo 介绍一下struct{}
		fmt.Println("收到消息，不买蛋了")
		return
	default:
	}
	fmt.Println("买蛋")
	go buySEgg(ctx)
	go buyBEgg(ctx)
	time.Sleep(1 * time.Second)
}

func buySEgg(ctx context.Context) {
	fmt.Println("去买蛋")
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done(): // todo 介绍一下struct{}
		fmt.Println("收到消息，不买蛋了buySEgg")
		return
	default:
	}
	fmt.Println("买蛋")
}
func buyBEgg(ctx context.Context) {
	fmt.Println("去买蛋")
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done(): // todo 介绍一下struct{}
		fmt.Println("收到消息，不买蛋了buyBEgg")
		return
	default:
	}
	fmt.Println("买蛋")
}
