package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func GetFibonacciSeries(n int) ([]int, error) {
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome!"))
	})
	http.HandleFunc("/fb", func(w http.ResponseWriter, r *http.Request) {
		var fabs []int
		for i := 0; i < 100000; i++ {
			fabs, _ = GetFibonacciSeries(50)
		}
		w.Write([]byte(fmt.Sprintf("%v", fabs)))
	})
	_ = http.ListenAndServe(":8080", nil)
}

/*

访问 http://localhost:8080/debug/pprof/  可看到一些分析

点击profile会进行30秒的采样，并下载文件。

也能用命令行

go tool pprof http://<host>:<port>/debug/pprof/profile?seconds=10

然后使用 top命令  （top排序,如按cum排序： top -cum）
命令： list 方法名
等。

*/

/*

常见的分析指标

wall time : 缓冲时间。 程序运行的绝对时间之间可能有阻塞
cpu time
block time
memory allocation
GC times/time spent



*/
