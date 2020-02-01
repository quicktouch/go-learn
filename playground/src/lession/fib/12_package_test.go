package fib

import (
	"lession/series"
	// 可以起一个别名
	cm "src/github.com/easierway/concurrent_map"
	"testing"
)

/**
1. 基本复用模块单元：以首字母大写来约定可以被包外代码访问
2. 代码的package可以和所在的目录不一致
3. 同一目录的go代码的package要保持一致


package

1. 通过go get来获取远程依赖   go get -u 强制从网路更新远程依赖
2. 注意代码在github上的组织形式，以适应go get。 直接从代码路径开始，不要有src

go path的问题:

1. 可以在ide中设置GOPATH的，但是可能会导致在ide外调试时不一致、 某些ide可能会有些问题
2. 作者建议使用.bash_profile(macos)中设置

（笔者用的goland，直接从GOPATH中设置PROJECT GOPATH 为当前的项目目录。
  Global GOPATH 设置为默认安装目录/User/用户名/go）
*/

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSeries(10))
	//t.Log(series.getFibonacciSeries(10))
}

/**
init 方法

1. main被执行前，所有依赖的package的init方法都会被执行
2. 不同包的init函数按照包导入的依赖关系决定执行顺序
3. 每个包可以有多个init函数
4. 每个源文件可有多个init函数
*/

/**
使用远程的package

使用远程的package  https://github.com/easierway/concurrent_map.git
go get -u github.com/easierway/concurrent_map   （发现目录多了 pkg 、src 目录）

自己的代码提交到github并且适应go get：
 直接以代码路径开始，不要有src
*/
func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}

/**
Go存在的依赖问题
1. 同一个环境下，不同项目使用同一包的不同版本  （因为go get下来会放到同一个go目录）
2. 无法管理对包的特定版本的依赖

为了解决这些问题
随着Go 1.5 release版本发布，vendor目录被添加到除了GoPATH和GOROOT之外的依赖目录查找的解决方案。
在Go 1.6之前，需要手动设置环境变量。

查找依赖包路径的优先级如下：
1. 当前包下的vendor目录
2. 想上级目录查找，知道找到src下的vendor目录
3. 在GOPATH下面查找依赖包
4. 在GOROOT目录下查找

常用的依赖管理工具
1. godep
2. glide     `brew install glide`
3. dep


glide说明
1. 安装 brew install glide
2. 进入目录，执行glide init
3. glide install
*/
