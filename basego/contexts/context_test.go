package contexts

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	ctx := context.Background()
	parent := context.WithValue(ctx, "key1", "123")
	sub := context.WithValue(ctx, "ke2", "34")
	fmt.Printf("%v\n", parent.Value("key1"))
	fmt.Printf("%v\n", sub.Value("ke2"))
}

func TestContext_Timeout(t *testing.T) {
	bg := context.Background()
	timeoutCtx, cancel1 := context.WithTimeout(bg, time.Second)
	subCtx, cancel2 := context.WithTimeout(timeoutCtx, 3*time.Second)
	go func() {
		<-subCtx.Done()
		fmt.Printf("timeout")
	}()
	time.Sleep(2 * time.Second)
	cancel2()
	cancel1()
}

func TestTimeoutExample(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	bsChan := make(chan struct{})
	go func() {
		slowBusiness()
		bsChan <- struct{}{}
	}()
	select {
	case <-ctx.Done():
		fmt.Println("timeout")
	case <-bsChan:
		fmt.Println("business end")
	}
}

func slowBusiness() {
	time.Sleep(2 * time.Second)
}

func TestTimeoutTimeAfter(t *testing.T) {
	bsChan := make(chan struct{})
	go func() {
		slowBusiness()
		bsChan <- struct{}{}
	}()

	timer := time.AfterFunc(time.Second, func() {
		fmt.Println("timeout")
	})
	<-bsChan
	fmt.Println("business end")
	timer.Stop()
}
