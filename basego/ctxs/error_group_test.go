package ctxs

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync/atomic"
	"testing"
	"time"
)

func TestErrgroup(t *testing.T) {
	eg, ctx := errgroup.WithContext(context.Background())

	var result int64 = 0

	for i := 0; i < 10; i++ {
		delta := i
		eg.Go(func() error {
			atomic.AddInt64(&result, int64(delta))
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
	ctx.Err()
	fmt.Println(result)
}

func TestBusinessTimeout(t *testing.T) {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	end := make(chan struct{}, 1)
	go func() {
		MyBusiness()
		end <- struct{}{}
	}()
	ch := timeoutCtx.Done()
	select {
	case <-ch:
		fmt.Println("timeout")
	case <-end:
		fmt.Println("business end")
	}
}

func TestParentValueCtx(t *testing.T) {
	ctx := context.Background()
	childCtx := context.WithValue(ctx, "map", map[string]string{})
	ccChild := context.WithValue(childCtx, "key1", "value1")
	m := ccChild.Value("map").(map[string]string)
	m["key1"] = "val1"
	val := childCtx.Value("key1")
	fmt.Println(val)
	val = childCtx.Value("map")
	fmt.Println(val)
}
func MyBusiness() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Hello world")
}

func TestParentCtx(t *testing.T) {
	ctx := context.Background()
	dlCtx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Minute))
	childCtx := context.WithValue(dlCtx, "key", 123)
	cancel()
	res := childCtx.Value("key")
	err := childCtx.Err()
	fmt.Println(err, res)
}

func TestContext(t *testing.T) {
	ctx := context.Background()
	valCtx := context.WithValue(ctx, "abc", 123)
	val := valCtx.Value("abc")
	fmt.Println(val)
}
