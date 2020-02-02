package main

import (
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

const (
	col = 10000
	row = 10000
)

func fillMatrix(m *[row][col]int) {
	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			m[i][j] = s.Intn(100000)
		}
	}
}

func calculate(m *[row][col]int) {
	for i := 0; i < row; i++ {
		tmp := 0
		for j := 0; j < col; j++ {
			tmp += m[i][j]
		}
	}
}

func main() {
	//创建输出文件
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}

	// 获取系统信息
	if err := pprof.StartCPUProfile(f); err != nil { //监控cpu
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	// 主逻辑区，进行一些简单的代码运算
	x := [row][col]int{}
	fillMatrix(&x)
	calculate(&x)

	f1, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	//runtime.GC()                                       // GC，获取最新的数据信息
	if err := pprof.WriteHeapProfile(f1); err != nil { // 写入内存信息
		log.Fatal("could not write memory profile: ", err)
	}
	_ = f1.Close()

	f2, err := os.Create("goroutine.prof")
	if err != nil {
		log.Fatal("could not create groutine profile: ", err)
	}

	if gProf := pprof.Lookup("goroutine"); gProf == nil {
		log.Fatal("could not write groutine profile: ")
	} else {
		_ = gProf.WriteTo(f2, 0)
	}
	_ = f2.Close()
}

// Go 支持的多种Profile
/*


go help testflag
https://golang.org/pkg/runtime/pprof/

```bash
# 编译测试代码
go build prof.go
# 运行二进制, 运行完多了 cpu.prof goroutine.prof mem.prof 这几个文件
./prof
```

# 查看这些文件  (适合短时间运行的程序)
go tool pprof  编译出的二进制  cpu.prof/goroutine.prof/mem.prof

例如:
```bash
go tool pprof prof cpu.prof

# log 如下,
# 随后使用top命令可以看到cpu详情。  flat表示所占时间和所占比例。 cum/cum%表示这个函数还调用了别的函数，总体加和在一起所占的时间和比例。
# 使用list + 符号名 命令可看具体哪行的耗时
# svg 以图像的方式查看调用关系及具体的cpu耗时，比较直观。
# exit 退出

File: prof
Type: cpu
Time: Feb 2, 2020 at 10:44pm (CST)
Duration: 2.06s, Total samples = 1.81s (87.91%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 1.81s, 100% of 1.81s total
Showing top 10 nodes out of 12
      flat  flat%   sum%        cum   cum%
     1.77s 97.79% 97.79%      1.78s 98.34%  main.fillMatrix
     0.02s  1.10% 98.90%      0.02s  1.10%  main.calculate
     0.01s  0.55% 99.45%      0.01s  0.55%  math/rand.(*rngSource).Int63
     0.01s  0.55%   100%      0.01s  0.55%  runtime.newstack
         0     0%   100%      1.81s   100%  main.main
         0     0%   100%      0.01s  0.55%  math/rand.(*Rand).Int31
         0     0%   100%      0.01s  0.55%  math/rand.(*Rand).Int31n
         0     0%   100%      0.01s  0.55%  math/rand.(*Rand).Int63
         0     0%   100%      0.01s  0.55%  math/rand.(*Rand).Intn
         0     0%   100%      0.01s  0.55%  os.Create
(pprof) list fillMatrix
Total: 1.81s
ROUTINE ======================== main.fillMatrix in /Users/panda/Desktop/go-learn/playground/src/lession/tools/file/prof.go
     1.77s      1.78s (flat, cum) 98.34% of Total
         .          .     15:
         .          .     16:func fillMatrix(m *[row][col]int) {
         .          .     17:   s := rand.New(rand.NewSource(time.Now().UnixNano()))
         .          .     18:   for i := 0; i < row; i++ {
         .          .     19:           for j := 0; j < col; j++ {
     1.77s      1.78s     20:                   m[i][j] = s.Intn(100000)
         .          .     21:           }
         .          .     22:   }
         .          .     23:}
         .          .     24:
         .          .     25:func calculate(m *[row][col]int) {
(pprof) exit
```
# go-torch

go-torch查看火炬图。

```bash
go-torch cpu.prof
```
*/

/**
以http的方式输出profile

- 简单，适合持续性运行的程序
- 应用程序中导入`import _ "net/http/pprof"`, 并启动http server 即可
- http://<host>:<port>/debug/pprof/
- go tool pprof http://<host>:<port>/debug/pprof/profile?seconds-10 (默认30s)
- go-torch -seconds 10 http://<host>:<port>/debug/pprof/profile

*/
