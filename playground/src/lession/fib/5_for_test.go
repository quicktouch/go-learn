package fib

import (
	"fmt"
	"runtime"
	"testing"
)

func TestFor(t *testing.T) {
	// while条件循环 条件 n < 5
	n := 0
	for n < 5 {
		n++
		t.Log(n)
	}
	t.Log("----")
	//无限循环
	n = 0
	for {
		n++
		if n > 5 {
			break
		}
		t.Log(n)
	}
}

// if 支持变量赋值
func TestIf(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("true")
	}
	//if v,err := someFun(); err = nil {
	//}else {
	//}
}

//switch
func TestSwitch(t *testing.T) {
	// 1. 不限制常量或者整数
	// 2. 不需要case break
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s", os)
	}

	// 3. case 表达式
	num := 99
	switch {
	case 0 <= num && num <= 88:
		t.Log("A")
	default:
		t.Log("B")
	}

	switch num {
	case 0, 2:
		t.Log("B3")
	case 100, 99:
		t.Log("B2")
	default:
		t.Log("B1")
	}
}
