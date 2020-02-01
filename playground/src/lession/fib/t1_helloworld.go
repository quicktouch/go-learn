package fib

//main包比较特殊。它定义了一个独立可执行的程序，而不是一个库。

import (
	"fmt"
	"os"
)

// 在main里的main 函数 也很特殊，它是整个程序执行时的入口
func main1() {
	fmt.Println("hello world")
	fmt.Println(len(os.Args))

	//读取命令的入参
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = ""
	}
	fmt.Println(s)

	//s := "" 		短变量声明，最简洁，但只能用在函数内部，而不能用于包变量
	//var s string  形式依赖于字符串的默认初始化零值机制，被初始化为""
	//var s = ""    第三种形式用得很少，除非同时声明多个变量
	//var s string = "" 第四种形式显式地标明变量的类型，当变量类型与初值类型相同时，类型冗余，但如果两者类型不同，变量类型就必须了。
	//实践中一般使用前两种形式中的某个，初始值重要的话就显式地指定变量的类型，否则使用隐式初始化。

}

//Go是一门编译型语言，Go语言的工具链将源代码及其依赖转换成计算机的机器指令（译注：静态编译）。
// Go语言提供的工具都通过一个单独的命令go调用，go命令有一系列子命令。
// 最简单的一个子命令就是run。这个命令编译一个或多个以.go结尾的源文件，链接库文件，并运行最终生成的可执行文件。（本书使用$表示命令行提示符。）

// 直接运行  go run
// 例如:  go run t1_helloworld.go  (需要将package main)

// 变异成可执行的二进制文件译注：Windows系统下生成的可执行文件是helloworld.exe，增加了.exe后缀名）
// go build t1_helloworld.go   可以看到目录下多了一个 t1_helloworld
// 可以直接执行 ./t1_helloworld

// Go语言的代码通过包（package）组织，
//包类似于其它语言里的库（libraries）或者模块（modules）。
//一个包由位于单个目录下的一个或多个.go源代码文件组成, 目录定义包的作用。
//每个源文件都以一条package声明语句开始，这个例子里就是package main, 表示该文件属于哪个包，紧跟着一系列导入（import）的包，之后是存储在这个文件里的程序语句。
