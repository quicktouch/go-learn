package fib

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
	"unsafe"
)

/*
常见的并发任务

- 仅执行一次 (单例模式：懒汉式，线程安全)

sync.Once 的 Do方法

```go
var once sync.Once
var obj *SingletonObj

func GetSingletonObj() * SingletonObj {
   once.Do(func(){
	  obj = &SingletonObj{}
   })
   return obj
}
```
*/

type Singleton struct {
}

var singleInstance *Singleton
var once sync.Once

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestOnce(t *testing.T) {
	var xx = GetSingletonObj()
	var xx1 = GetSingletonObj()
	t.Log(xx)
	t.Log(xx1)

	fmt.Printf("---%x\n", unsafe.Pointer(xx))
	fmt.Printf("---%x\n", unsafe.Pointer(xx1))
}

/*
仅需任意任务完成，使用channel 。
需要使用buffer否则会阻塞协程
*/

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("the result is from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	// 一旦有数据，不阻塞，直接执行。 随后会阻塞，等待消息被取走
	// buffer解耦
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine()) // 2
	t.Log(FirstResponse())
	time.Sleep(time.Second * 2)
	t.Log("After:", runtime.NumGoroutine()) // no buffer 11.  buffer 2
}

/**
所有的任务完成
*/

func myTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("the result is from %d", id)
}

func AllResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := myTask(i)
			ch <- ret
		}(i)
	}
	finalRet := ""
	for j := 0; j < numOfRunner; j++ {
		finalRet += <-ch + "\n"
	}
	return finalRet
}

func TestAllResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine()) // 2
	t.Log("xxx:", AllResponse(), "xxx\n")
	time.Sleep(time.Second * 2)
	t.Log("After:", runtime.NumGoroutine()) // no buffer 11.  buffer 2
}

/*
对象池

经常遇见，例如数据库连接，网络连接。池化，避免重复创建。

GO： 使用buffered channel实现对象池
*/

type ReusableObj struct {
}

type ObjPool struct {
	bufChan chan *ReusableObj
}

func NewObjPool(numObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ReusableObj, numObj)
	for i := 0; i < 10; i++ {
		objPool.bufChan <- &ReusableObj{}
	}
	return &objPool
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout): // 超时控制
		return nil, errors.New("time out")
	}
}

func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	// size已经满了的情况是放不进去的
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}

func TestPool(t *testing.T) {
	pool := NewObjPool(10)
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("v:%T \n", v)
			if err := pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
		}
	}
	fmt.Printf("Done")
}

/*
sync.pool 对象缓存
- 尝试从私有对象获取
- 私有对象不存在，尝试从当前的Processor的共享池获取
- 如果当前Processor共享池也是空的，那么尝试从Processor的共享池获取
- 如果所有的子池都是空的，最后就用用户指定的New函数产生一个新的对象返回

![20200131158045064229745.png](http://image.tyrad.cc/20200131158045064229745.png)

对象的放回
- 如果私有对象不存在则保存为私有对象
- 如果私有对象存在，放入当前Processor子池的共享池中

声明周期
- GC会清除sync.pool缓存的对象
- 对象的缓存有效期为下一次GC之前

总结
- 适用于通过复用，降低复杂对象的创建和GC
- 协程安全，会有锁的开销
- 声明周期受GC影响，不适合于做连接池等需要自己管理声明周期的资源的池化
*/

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create a new project.")
			return 100
		},
	}
	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)
	//runtime.GC() //GC会清除sync.pool中的缓存对象    调用后v1 = 100 （重新创建）
	v1, _ := pool.Get().(int)
	fmt.Println(v1) // 3
}
