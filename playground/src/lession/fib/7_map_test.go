package fib

import "testing"

func TestMap(t *testing.T) {

	// map 声明
	m := map[string]int{"one": 1, "two": 2}
	m1 := map[string]int{}
	t.Log(m)  // map[one:1 two:2]
	t.Log(m1) // map[]

	m2 := make(map[string]int, 10 /*inital capcity*/)
	t.Log(m2)

	m2["11"] = 16
	t.Log(m2)      // map[11:16]
	t.Log(len(m2)) // 1
}

func TestAccessNotExistingKey(t *testing.T) {
// 初始化空map
m1 := map[int]int{}
// 不存在的key,返回零值. 不会返回nil
t.Log(m1[1]) // 0

m1[3] = 100
// 因此我们取值的时候需要判断是否存在
if v, ok := m1[3]; ok {
	t.Log(ok, ". value is ", v) // true . value is  100
} else {
	t.Log("key 3 not  existing.")
}
}

func TestTravelMao(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2}
	for k, v := range m {
		t.Log(k, v)
	}
}

// map 与工厂模式
//1. map的value可以是一个方法
//2. 与Go的Dock type接口方式一起，可以方便的实现单一方法对象的工厂模式

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int {
		return 10 * op
	}
	m[2] = func(op int) int {
		return op
	}
	t.Log(m, m[1](1), m[2](2)) //  map[1:0x10faf80 2:0x10fafa0]   10   2
}

// 实现Set
// 没有内置的Set ,可以用map[type]bool
// 1. 元素唯一性
// 2. 基本操作  a.添加 b.判断元素是否存在  c.删除元素 d.元素个数
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 1
	if mySet[n] {
		t.Logf("%d is exsisting", n)
	} else {
		t.Logf("%d not exsisting", n)
	}
	//删除元素
	delete(mySet, 1)

}
