package fib

import (
	"fmt"
)

// 声明语句:
// var    变量
// const  常量
// type   类型
// func   函数

// 函数的结构:
// 1. 先声明package属于哪一个包
// 2. 使用import导入依赖的包
// 3. 包一级的类型、变量、常量、函数的声明语句 (顺序无关，但是要先声明再使用)

//------------------------------------------------------------------

// 包一级范围声明语句声明
// 包一级声明语句声明的名字可在整个包对应的每个源文件中访问
const boilingF = 212.0

func mainxx() {
	// f c 为函数内部变量
	var f = boilingF
	var c = fToC(f)
	fmt.Printf("boiling point = %g° or %g摄氏度 \n", f, c)

	varsDebug()
	p()
	newX()
	setValue()

	c1 := FtoC(212.0)
	fmt.Println(c1.String())
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

// 变量
// var 变量名字 [类型]  [= 表达式]
// 其中 [类型]和[= 表达式] 两者可以省略其中一个
/// 如果省略的是类型: 可以根据初始化表达式自动推导类型
/// 如果省略的是[= 表达式]: 用零值初始化该变量。
//			数值类型变量对应的零值是0，
//			布尔类型变量对应的零值是false，
//			字符串类型对应的零值是空字符串，
//			接口或引用类型（包括slice、指针、map、chan和函数）变量对应的零值是nil
//			数组或结构体等聚合类型对应的零值是每个元素或字段都是对应该类型的零值。

func varsDebug() {
	//初始化后自动设置为 ""
	var s string
	fmt.Printf(s + "\n")

	//声明一组变量
	var i, j, k int
	fmt.Println(i, j, k)

	//初始化一组变量
	var b, f, t = true, 2.3, "four"
	fmt.Println(b, f, t)

	//通过使用函数的返回值，返回变量
	//m, err := os.Open("")
	//简短变量声明语句中必须至少要声明一个新的变量,否则不能编译通过
	//m, err := os.Open("")
	//if err != nil {
	//	return err
	//}
	////...
	//f.Close()
}

/// 指针
/// 一个变量对应一个保存了变量对应类型值的内存空间。
/// 普通变量在声明语句创建时被绑定到一个变量名,比如叫x的变量,但是还有很多变量始终以表达式方式引入，例如x[i]或x.f变量。
//  所有这些表达式一般都是读取一个变量的值，不过位于等号左边的话就是赋值。
// 指针的值是变量的地址（内存的存储位置）。 不是每一个值都有变量地址，但是每一个变量必然有对应的内存地址。
// 通过地址可以读取或者更新变量的值，而不需要知道变量的名字。
func p() {
	x := 111
	p := &x         // points to x   &为取地址符
	fmt.Println(*p) // 111
	*p = 2
	fmt.Println(x) //  2

	// 每次的地址都不同，返回false
	fmt.Println(p1() == p1())
	v := 1
	incr(&v)              // 2
	fmt.Println(incr(&v)) // 3
}

// 可以函数后返回局部变量地址. 也是安全的.
//在局部变量地址被返回之后依然有效，因为指针p依然引用这个变量。
//var p = f()
func p1() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++ // 增加p的指针变量的值。不会改变指针
	return *p
}

// new函数
// 表达式new(T)将创建一个T类型的匿名变量，初始化为T类型的零值，然后返回变量地址，返回的指针类型为*T
func newX() {
	p := new(int)   // p, *int 类型, 指向匿名的 int 变量
	fmt.Println(p)  // 地址
	fmt.Println(*p) // 0
}

// 变量的生命周期
// 生命周期指的是在程序运行期间变量有效存在的时间间隔。
// 对于在包一级声明的变量来说，它们的生命周期和整个程序的运行周期是一致的。
// 而相比之下，局部变量的声明周期则是动态的：每次从创建一个新变量的声明语句开始，直到该变量不再被引用为止，然后变量的存储空间可能被回收。
// 函数的参数变量和返回值变量都是局部变量。它们在函数每次被调用的时候创建。

// 编译器会自动选择在栈上还是在堆上分配局部变量的存储空间，
// 但可能令人惊讶的是，这个选择并不是由用var还是new声明变量的方式决定的。

// 例如: 下面
// x变量必须在堆上分配，因为它在函数退出后依然可以通过包一级的global变量找到，
// 虽然它是在函数内部定义的；用Go语言的术语说，这个x局部变量从函数f中逃逸了。
// 逃逸的变量需要额外分配内存，同时对性能的优化可能会产生细微的影响

//var global *int
//
//func f() {
//	var x int
//	x = 1
//	global = &x
//}
//
//func g() {
//	y := new(int)
//	*y = 1
//}

// 元组赋值
// x,y = 1,2
// 可赋值性
func setValue() {
	medals := []string{"gold", "sliver", "bronze"}
	medals[0] = "golden"
	fmt.Println(medals)
}

// 类型
// 可以声明新的类型名称,如下
// type 类型名称 底层类型
// 不同类型的进行比较，虽然低层相同但是声明的类型不同，不能比较 `compile error: type mismatch`

type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

func CToF(c Celsius) Fahrenheit {
	// 需要显式转型操作才能将float64转为对应的类型
	return Fahrenheit(c*9/5 + 32)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// 下面的声明语句，Celsius类型的参数c出现在了函数名的前面，
//表示声明的是Celsius类型的一个名叫String的方法，该方法返回该类型对象c带着°C温度单位的字符串：
//c1 := FtoC(212.0)
//fmt.Println(c1.String())
// 相当于添加了扩展
func (c Celsius) String() string {
	return fmt.Sprintf("%g摄氏度", c)
}

// 包和文件
// 可以通过包名.xx 的形式访问

// 包的初始化
// 1. 包的初始化首先是解决包级变量的依赖顺序，然后按照包级变量声明出现的顺序依次初始化：
//var a1 = b + c // a 第三个初始化, 为 3
//var b = f()   // b 第二个初始化, 为 2, 通过调用 f (依赖c)
//var c = 1     // c 第一个初始化, 为 1
//
//func f() int { return c + 1 }

// 2. 如果有多个.go源文件。 如果有初始化表达式则用初始化表达式，还有一些没有初始化表达式的，
// 例如某些表格数据初始化并不是一个简单的赋值过程。在这种情况下，
// 我们可以用一个特殊的init初始化函数来简化初始化工作。每个文件都可以包含多个init初始化函数
// func init() {}

// init方法除了不能被调用或引用外，其他行为和普通函数类似。
// 每个包在解决依赖的前提下，以导入声明的顺序初始化，每个包只会被初始化一次。
// 因此，如果一个p包导入了q包，那么在p包初始化的时候可以认为q包必然已经初始化过了。
// 初始化工作是自下而上进行的，main包最后被初始化。
// 以这种方式，可以确保在main函数执行之前，所有依赖的包都已经完成初始化工作了
