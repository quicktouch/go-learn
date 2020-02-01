package fib

import "testing"

/**
bool
string
int unit  int64 int32 等
byte
rune
float32 float64
complex64 complex128

1. 不支持隐式类型转换。
*/
type MyInt int

func TestImplicit(t *testing.T) {
	var a MyInt = 1
	var b int = 111
	//t.Log(a+b) 编译错误
	t.Log(a + MyInt(b))
}

//2. 不支持指针运算
//// string是值类型，其默认的初始化值为空字符串，而不是nil
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr += 1 编译错误，不支持指针运算
	t.Log(a, aPtr)
	// %T 打印类型
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log(s)      // "
	t.Log(len(s)) // 0
	//字符串不会为nil
}
