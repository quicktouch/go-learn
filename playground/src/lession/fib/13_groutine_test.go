package fib

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
Thead  vs Groutine

1). 创建时默认的stack（栈）的大小
- JDK5以后默认Java Thread stack 默认是1M
- Groutine的Stack初始化大小为2k,创建起来也更快

2). 和 系统线程KSE (Kernel Space Entity)的对应关系
- java thread是1：1
- Groutine 是M:N 多对多

如果thread和KSE是1：1关系的话。KSE是CPU直接调度，效率很高，但是如果线程之间发生切换，会牵扯到内核对象的切换，开销很大。
如果多个协程都在和一个内核对应，那么久减少了开销，go就是这么做的。

![](http://image.tyrad.cc/20200130158037814738545.png)


P为go语言的携程处理器。
P都在系统线程里，每个P都挂着一些携程队列G。有的是正在运行状态的，如G0.

如果有一个G非常占用时间，那么队列的G会不会被延迟很久？
不会。Go有守护线程，进行计数，计算每个P完成G的数量，如果某一个P一段时间完成的数量没有变化。
就会往携程的任务栈里插入特别的标记，当G运行的时候遇到非内联函数（？）就会读到这个标记，把自己中断，放到队尾。

？另外一个机制，当某个G被系统中断了，io需要等待时（？），P将自己移到另一个可使用的线程中，继续执行他的队列的G.


![](http://image.tyrad.cc/20200130158037816680653.png)

*/

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		// Go的方法调用传递时都是值传递，i被复制了一份，地址是不同的。因此可以执行
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Millisecond * 50)
	// 可以看到打印是乱序的 (调用程序一定要用function(){}(参数))
}

/**
共享内存并发机制。

其他语言：使用锁进行并发控制。

Lock lock = ...
lock.lock();
try {
  //process (thread-safe)
}catch(e){
}
finally{
	lock.unlock()
}

go:

package sync
Mutex   可lock和unlock
RWLock  对读锁和写锁进行分开。 当共享的内存被读锁锁住时，另外一个读锁去锁它时候，不是互斥锁。当写锁遇到它的时候才是互斥的。 比完全互斥锁的情况效率高一些。
*/

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Log(counter) // 线程不安全 结果不是5000
}

func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Log(counter) // 线程安全 结果5000
}

func TestWaitGroup(t *testing.T) {
	// 等待线程完成在往下执行。
	// 只有当wait的完成完后才往下执行
	/**
	var wg sync.WaitGroup
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
			//do some thing
		}()
	}
	wg.Wait()
	*/
	var wg sync.WaitGroup
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		//增加等待的量
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			//告知任务结束
			wg.Done()
		}()
	}
	//等待都执行完成
	wg.Wait()
	t.Log(counter) // 线程安全 结果5000
}

/**
CPS （communicating sequential process）并发机制
依赖通道来完成两个实体间的通信。

![20200130158038447911577.png](http://image.tyrad.cc/20200130158038447911577.png)
![20200130158038448661074.png](http://image.tyrad.cc/20200130158038448661074.png)

和Actor的直接通信不同，gsp模式则是通过channel进行通信的，更松散一些。
Go中的channel是有容量限制并且独立于处理Groutine ,而如erlang，actor模式中的mailbox容量是无限的，接收进行总是被动的处理消息。

channel分为两种模式，如下图:
![20200130158038496929185.png](http://image.tyrad.cc/20200130158038496929185.png)

方式1. 通信的两端必须同时在chanel上，一方不在的话就会进行等待阻塞。
方式2. bufferChannel 容量有限制
*/

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("---1s working on something else ")
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("---1s Task is done.")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
	/*
			=== RUN   TestService
			Done
			---1s working on something else
			---1s Task is done.
			--- PASS: TestService (1.06s)
			PASS
		可以发现他们是串行的
	*/
}

/*
实现类似java futureService的运行方式
*/
func AsyncService() chan string {
	// 没有使用buffer
	retCh := make(chan string)
	// 使用buffer的情况
	//retCh := make(chan string, 1)
	// 调用时启动另外一个协程，不阻塞当前协程
	go func() {
		ret := service() // "Done" 50ms
		fmt.Println("return result.")
		//运行完将结果放到channel
		retCh <- ret
		//是否计算完成并将结果返回给了channel，是否就退出，并做下一步的处理？
		//否。 没有使用buffer. 实际上会被阻塞
		fmt.Println("service exited.")
	}()
	// 返回channel，在channel上等待
	return retCh
}

func AsyncBufferService() chan string {
	// 使用buffer的情况
	retCh := make(chan string, 1)
	// 调用时启动另外一个协程，不阻塞当前协程
	go func() {
		ret := service() // "Done" 50ms
		fmt.Println("return result.")
		//运行完将结果放到channel
		retCh <- ret
		//是否计算完成并将结果返回给了channel，是否就退出，并做下一步的处理？
		//使用了buffer，buffer内会直接执行
		fmt.Println("service exited.")
	}()
	// 返回channel，在channel上等待
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()          // 1 s
	fmt.Println(<-retCh) //从channel取数据
	/*
		=== RUN   TestAsyncService
		---1s working on something else
		return result.
		---1s Task is done.
		Done
		service exited.
		--- PASS: TestAsyncService (1.00s)
		PASS
	*/

	retChBuffer := AsyncBufferService()
	otherTask()
	fmt.Println(<-retChBuffer)

	/*
		=== RUN   TestAsyncService
		---1s working on something else
		return result.
		service exited.
		---1s Task is done.
		Done
		--- PASS: TestAsyncService (1.00s)
	*/
}

/*
多路选择和超时

![20200130158038986977778.png](http://image.tyrad.cc/20200130158038986977778.png)

多渠道选择:
任何一个不阻塞，就会执行case语句，和case的写的顺序无关。
如果都没准备好，直接`default`

超时的控制：
time.After(...)
*/

func TestSelect(t *testing.T) {
	ret := <-AsyncService() // 50ms
	t.Log(ret)
}

func TestTimeoutSelect(t *testing.T) {
	select {
	case ret := <-AsyncService(): // 50ms
		t.Log(ret)
	case <-time.After(time.Millisecond * 30):
		t.Error("time out")
	}
}

/*
chanel的关闭和广播
*/

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		//Receiver可能不知道Producer什么时候结束。
		//引出关闭channel
		close(ch)

		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			data := <-ch
			fmt.Println(data)

			//1. 向关闭的channel发送数据，会导致panic
			//2. data, ok := <-ch , ok =false表示通道被关闭
			//3. 所有的channel接受者会在channel关闭时，立刻从阻塞等待中返回上述ok值为false。
			//   这个广播机制常常被利用,进行多个订阅者同时发送信号。 如退出信号

			//如何知道channel关闭.  ok =false表示通道被关闭
			if data, ok := <-ch; ok {
				fmt.Println(data)
			}
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()

	/*
		=== RUN   TestCloseChannel
		0
		1
		2
		3
		4
		5
		6
		7
		8
		9
		--- PASS: TestCloseChannel (0.00s)
		PASS
	*/
}

/**
任务的取消
*/

func isCancelled(cancelChan chan struct{}) bool {
	select {
	// 检查是否收到了消息。
	case <-cancelChan:
		return true
		// 阻塞中表示没有取消
	default:
		return false
	}
}

// 第一种cancel实现
func cancel_1(cancelChan chan struct{}) {
	// 发送空的
	cancelChan <- struct{}{}
}

// 第二种cancel的实现
func cancel_2(cancelChan chan struct{}) {
	close(cancelChan)
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Canceled")
		}(i, cancelChan)
	}
	//cancel_1(cancelChan) // 4 Canceled
	cancel_2(cancelChan) // 广播机制。 均触发 canceled
	time.Sleep(time.Second * 1)
}

