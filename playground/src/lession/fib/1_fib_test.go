package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	//写法1: 通常用于声明变量
	//var a int = 1
	//var b int = 2
	// 写法2: 通常用于声明变量
	//var (
	//	a int = 1
	//	b     = 2
	//)
	//写法3
	a := 1
	b := 2
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	// 可以在一个赋值语句中对多个语句进行赋值
	a,b = b,a
	t.Log(b,a)
}

// 常量