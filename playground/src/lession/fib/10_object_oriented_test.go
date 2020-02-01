package fib

import (
	"fmt"
	"testing"
	"unsafe"
)

/**
go是面向对象的语言吗？
官方回答：Yes and no

go不支持继承
*/

// 结构体定义
type Employee struct {
	Id   string
	Name string
	Age  int
}

// 行为和方法的定义，两种方法

// 1. 实例对应方法被调用时，实例的成员会进行值复制
//func (e Employee) String() string {
//	return fmt.Sprintf("id:%s name:%s age:%d", e.Id, e.Name, e.Age)
//}

// 2. 通常为了避免内存拷贝我们使用第二种定义方式
func (e Employee) String() string {
	fmt.Printf("address is %x\n", unsafe.Pointer(&e.Name))
	// 通过指针访问，直接.就行
	return fmt.Sprintf("id:%s name:%s age:%d", e.Id, e.Name, e.Age)
}

func TestObjectOriented(t *testing.T) {
	// struct初始化的方法
	e := Employee{"0", "Bob", 20}

	e1 := Employee{Name: "Mike", Age: 30}

	e11 := Employee{}

	e2 := new(Employee) //返回指针 相当于 &Employee{}
	e2.Id = "22"
	e2.Age = 22
	e2.Name = "xiaoming"

	t.Log(e)   //{0 Bob 20}
	t.Log(e1)  // { Mike 30}
	t.Log(e11) // {  0}
	t.Log(e2)  // &{22 xiaoming 22}
	t.Log(e.String())
	t.Log(e2.String())
}

func TestStructOperation(t *testing.T) {
	e := Employee{"1", "bob", 20}
	fmt.Printf("address1 is %x \n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}

/**
接口特性：
1. 接口是非侵入性的，实现不依赖接口定义
2. 所以接口的定义可以包含在接口使用者包内
*/

type Programmer interface {
	WriteHelloWorld() string
}
type GoProgrammer struct {
}

// "duck type"  方法签名看起来是一样的。
func (p *GoProgrammer) WriteHelloWorld() string {
	return "hello world"
}

func TestInterface(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}

// 接口变量

// 自定义类型
/**
 */

// go中实现继承和override
type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

// 匿名嵌套类型
/// 可以实现类似继承的效果，但是有所不同
/// 1. 重写(Override)

type Dog1 struct {
	Pet
}

// 进行方法的重写
func (p *Dog1) Speak() {
	fmt.Print("dog1 ...")
}

func (p *Dog1) SpeakTo(host string) {
	p.Speak()
	fmt.Println("dog1 ", host)
}

func TestDog(t *testing.T) {
	dog := new(Dog1)
	dog.SpeakTo("hhh")
	dog.Speak()
}

// 多态
// 多态就是同一个接口，使用不同的实例而执行不同操作，
// 如 interface 打印；  彩色相机和黑白相机执行效果不同
type MyProgrammer interface {
	sayHello() string
}
type Golang struct {
}
type Java struct {
}
type Ruby struct {
}

func (p *Golang) sayHello() string {
	return "go white hello world"
}

func (p *Java) sayHello() string {
	return "java white hello world"
}

func callSayHello(p MyProgrammer) {
	fmt.Println(p.sayHello())
}

func TestDuotai(t *testing.T) {
	goProg := new(Golang)
	javaProg := &Java{}
	// 需要操作指针
	callSayHello(goProg)
	callSayHello(javaProg)
	//callSayHello(new(Ruby))  因为没有实现方法，这里编译报错
}

/// 空接口和断言
/**
1. 空接口可以表示任何类型  类似 void*   或 java的Object
2. 通过断言将空接口转换为定制类型： v,ok = p.(int)
*/

func DoSomething(p interface{}) {

	/**
	if i, ok := p.(int); ok {
		fmt.Println("int", i)
		return
	}
	if i, ok := p.(string); ok {
		fmt.Println("string", i)
		return
	}
	fmt.Println("Unknown Type")
	*/

	//可以将上述代码进行简化
	switch v := p.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("unknown")
	}
}

func TestConvert(t *testing.T) {
	DoSomething(10)
	DoSomething("aa")
	DoSomething(true)
	//int 10
	//string aa
	//unknown
}

/**
Go接口最佳实践:
1). 倾向于使用小的接口定义，很多接口只包含一个方法
2). 较大的接口定义，可以由多个小接口组合而成
```go
type ReadWrite interface {
	Reader
	Write
}
```
3). 只依赖必要功能的最小接口. 方便复用
```go
func StoreData(reader Reader) error{...}
```
*/
