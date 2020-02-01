package fib

import "testing"

// 常量

// 快速设置连续值

// 递增加一
const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
)

// 位运算
const (
	Open = 1 << iota
	Close
	Pending
)

func TestXx(test *testing.T) {
	test.Log(Monday, Tuesday, Wednesday)
	test.Log("-----")
	test.Log(Open, Close, Pending)
	test.Log(Pending | Close)
}
