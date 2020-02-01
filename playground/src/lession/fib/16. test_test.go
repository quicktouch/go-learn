package fib

import (
	"fmt"
	"testing"
)

/*
单元测试

1. 文件 _test结尾，方法名Test开头

- Fail,Error: 测试失败，该测试继续，其他测试也继续执行
- FailNow, Fatal: 该测试失败，该测试终止，其他测试继续执行

代码覆盖率
go test -v -cover
断言:
https://github.com/stretchr/testify
*/

func TestForSqure(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		if ret != expected[i] {
			t.Errorf("input %d expected %d, the actual is %d",
				inputs[i], expected[i], ret)
		}
	}
}

func square(i int) int {
	return i * i
}

func TestErrorInCode(t *testing.T) {
	fmt.Println("Start")
	t.Error("Error")
	fmt.Println("End")
}

func TestFailInCode(t *testing.T) {
	fmt.Println("Start")
	t.Fatal("Error")   //该测试失败，该测试终止，其他测试继续
	fmt.Println("End") //不会执行本行
}

/*
BenchNow 性能测评

运行所有文件 go test -bench=.   [-benchmenm (分析内存大小，allocs内存个数)]
运行方法  go test -bench=方法名 [-benchmenm (分析内存大小，allocs内存个数)]
*/
func BenchmarkXX(b *testing.B) {
	//性能无关的代码
	//....
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

	}
	b.StopTimer()
	//性能无关的代码
	//...
}

/*
BDD  (Behavior Driven Developement)

BDD in Go
地址： https://github.com/smartystreets/goconvey
安装：
 go get -u  github.com/smartystreets/goconvey

启动Web ui
 $GOPATH/bin/goconvey
*/
