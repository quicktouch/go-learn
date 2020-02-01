package fib

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestFuc(t *testing.T) {

	//函数是一等公民
	//差异
	/**
	1. 支持多个返回值
	2. 所有参数都是值传递： slice map channel会有传引用的错觉
	3. 函数可以作为变量的值
	4. 函数可以作为参数和返回值
	*/
	t.Log(returnMultiValues())

	// 函数运算计时
	timeSpent(slowFunc)(1)
	t.Log(Sum(1, 2, 3))    // 6
	t.Log(Sum(1, 2, 3, 4)) // 10
}

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

//函数一等公民（闭包）

func timeSpent(inner func(op int) int) func(op int) int {
	//入参是函数 返回也是函数
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

// 可变参数
func Sum(opts ...int) int {
	ret := 0
	for _, op := range opts {
		ret += op
	}
	return ret
}

// 延迟执行函数
func TestDefer(t *testing.T) {
	defer func() {
		t.Log("Clean resources")
	}()
	t.Log("Started")
	panic("Fatal error") //异常仍会执行defer
}
