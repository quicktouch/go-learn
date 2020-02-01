package fib

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// Context与任务取消

// 任务可能是关联的，任务又包括子任务。 1.9以后引入了context
/*
- 根context： 通过context.Background()创建
- 子context：context.WithCancel(parentContext)
	- ctx,cancel "= context.WithCancel(context.Background())
- 当前Context被取消时，基于它的子context都会取消
- 接收取消通知 <-ctx.Done()
*/

func isCanceled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCanceled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			// 上面的死循环被打破，才会打印
			fmt.Println(i, "Canceled")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}
