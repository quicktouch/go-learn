package fib

import "testing"

func TestArray(t *testing.T) {
	// 数组声明, 并赋值
	var a [3]int
	a[0] = 1 // 直接通过下标访问或修改
	t.Log(a) // [1 0 0]

	// 数组声明的同时进行赋值
	b := [3]int{1, 2, 3}
	c := [2][2]int{{1, 2}, {3, 4}}
	t.Log(b, c) // [1 2 3] [[1 2] [3 4]]
}

func TestArrayTravel(t *testing.T) {
	// 数组遍历，直接访问下标或者foreach
	arr3 := [...]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	t.Log("----")
	for index, e := range arr3 {
		t.Log(index, e)
	}
	t.Log("----")
	// 使用`_`忽略
	for _, e := range arr3 {
		t.Log(e)
	}
}

func TestArraySlice(t *testing.T) {
	//数组的截取操作
	a := [...]int{1, 2, 3, 4, 5, 6, 7}

	// a[开始索引(包含), 结束索引(不包含)]   =>  [index1,index2)
	t.Log(a[1:2])      // 2
	t.Log(a[1:len(a)]) //[2 3 4 5 6 7]
	t.Log(a[:])        //[1 2 3 4 5 6 7]

	t.Log("-----------")

	//切片
	//切片的数据结构： 指针 、长度、 容量
	var s0 []int
	t.Log(s0, len(s0), cap(s0))

	s0 = append(s0, 1)          // [] 0 0
	t.Log(s0, len(s0), cap(s0)) // [1] 1 1

	//可以使用make来初始化
	s2 := make([]int, 3, 5)     //长度是3 容量是5
	t.Log(s2, len(s2), cap(s2)) //[0 0 0] 3 5
	// t.Log(s2[4]) 会越界
	s2 = append(s2, 8)
	t.Log(s2, len(s2), cap(s2)) //[0 0 0 8] 4 5

	//切片growth ,长度可变
	//两种声明方法均可
	//var s1 []int
	s1 := []int{}
	for i := 0; i < 10; i++ {
		s1 = append(s1, i)
		t.Log("--", len(s1), cap(s1))
		//  -- 1 1
		//  -- 2 2
		//  -- 3 4
		//  -- 4 4
		//  -- 5 8
		//  -- 6 8
		//  -- 7 8
		//  -- 8 8
		//  -- 9 16
		//  -- 10 16
		//可以看到容量呈倍数x2增长
	}
	// append好处是多个切片可以共享一个结构
}
