package main

import (
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

const row = 100
const col = 100

func fillMatrix(m *[row][col]int) {
	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			m[i][j] = s.Intn(10000)
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
	// CPU
	//创建输出文件
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
		return
	}
	// 获取系统信息
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile:", err)
	}
	defer pprof.StopCPUProfile()

	// 测试代码
	x := [row][col]int{}
	fillMatrix(&x)
	calculate(&x)

	// 堆
	f1, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
	if err := pprof.WriteHeapProfile(f1); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
	_ = f1.Close()

	// goroutine
	f2, err := os.Create("goroutine.prof")
	if err != nil {
		log.Fatal("could not create goroutine profile", err)
	}
	if qProf := pprof.Lookup("qoroutine"); qProf == nil {
		log.Fatal("could not write goroutine profile: ", err)
	}
	_ = f2.Close()
}
