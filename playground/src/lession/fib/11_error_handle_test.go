package fib

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

/**

Go错误机制
1. 没有异常机制
2. error类型实现了error接口
3.可以通过errors.New来快速创建错误实例

type error interface {
	Error() string
}

func TestName(t *testing.T) {
	errors.New("xx")
}
*/
var LessThanTwoError = errors.New("n should not less than 2")
var LargerThan100Error = errors.New("n should not large than 100")

func GetFibonacci(n int) ([]int, error) {
	//及早失败，避免嵌套
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerThan100Error
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func GetFibonacci2(str string) {
	var (
		i    int
		err  error
		list []int
	)
	//及早失败，避免嵌套
	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("error", err)
		return
	}
	if list, err = GetFibonacci(i); err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(list)
}

func TestFab(t *testing.T) {
	if v, err := GetFibonacci(1000); err != nil {
		t.Log(err)
		if err == LessThanTwoError {
			t.Log("--2")
		}
		if err == LargerThan100Error {
			t.Log("--100")
		}
	} else {
		t.Log(v)
	}
	GetFibonacci2("1001")
}

/**
panic 和 recover

panic
1. panic用于不可以恢复的错误
2. panic退出前会执行defer指定的内容

os.Exit
1. 不会调用defer指定函数
2. 不会打印栈的信息
*/

func TestPanicVxExit(t *testing.T) {
	defer func() {
		// panic会执行defer的代码，而且打印调用栈的信息
		// os.Exit不会
		fmt.Println("finally")
	}()
	fmt.Println("start")
	panic(errors.New("Something error"))
	//os.Exit(-1)
}

/**
recover
可以类比 try{}catch{}
recover用于错误的恢复：

defer func() {
	if error := recover(); err != nil {
		//恢复错误
	}
}()

注意:  不要滥用
1. 形成僵尸服务进程，导致health check失效
2. "let it crash!"往往是我们恢复不确定性错误的好方法
*/

func TestPanicVxExit2(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			// 程序没有打印调用栈，正常退出
			fmt.Println("final:recover", err) // final:recover Something error
		}
	}()
	fmt.Println("start")
	panic(errors.New("Something error"))
}
