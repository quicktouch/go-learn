package fib

/*

安装graphviz  添加图形的支持

   brew install graphviz

安装go-torch

  1. go get github.com/uber/go-torch
  2. git clone https://github.com/brendangregg/FlameGraph.git 将 flamegraph.pl 拷贝到`/usr/local/bin`目录下
  3. 输入 `flamegraph.pl -h` 可以看到是否成功


## 通过文件方式输出 Profile

- 灵活性高，适用于特定代码段的分析
- 通过手动调用 runtime/pprof的api
- api相关文档 https://studygolang.com/static/pkgoc/pkg/runtime_pprof.htm
- go tool pprof [binary] [binary.prof]

可以输出cpu 内存 线程数



*/
