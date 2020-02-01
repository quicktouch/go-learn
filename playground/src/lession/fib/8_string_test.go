package fib

import (
	"strconv"
	"strings"
	"testing"
)

/**
和其他编程语言的差异
1. string是数据类型，不是引用或者指针类型
2. string是只读的byte slice， len函数返回其包含的byte数
3. string的byte数组可以存放任何数据
*/

func TestStrings(t *testing.T) {
	var s string
	t.Log(s) //默认值为 ""
	s = "hello world"
	t.Log(s, len(s))   //hello world  11
	s = "\xE4\xB8\xA5" // '严'字的二进制存储。 长度3个byte
	t.Log(s, len(s))   // 严   3

	//string是不可变的byte slice
	//s[1] = '3'   编译报错

	//Unicode UTF8
	//1. unicode是一种字符集（code point）
	//2. utf8是unicode的存储实现（转换为字节序列的规则）

	// 中  => unicode 4e2d  => string/[]byte  [0xe4,0xb8,0xad]

	s = "中"
	t.Log(len(s))                // 6
	c := []rune(s)               // rune表示它的unicode
	t.Log(c)                     // [20013 22269]
	t.Logf("中 unicode %x", c[0]) // 中 unicode 4e2d
	t.Logf("中 utf8 %x", s)       // 中 utf8 e4b8ad

	//遍历字符串的每个rune
	sx := "中国"
	for _, c := range sx {
		//log中 [1]表示与第一个参数对应
		t.Logf("%[1]c %[1]d %[1]x", c)
		// 中 20013 4e2d
		// 国 22269 56fd
	}

	// 字符串分割
	s1 := "A,B,C"
	parts := strings.Split(s1, ",") //[A B C]
	for i, part := range parts {
		t.Log(part, i)
	}
	//字符串拼接
	joinded := strings.Join(parts, "-")
	t.Logf(joinded) //A-B-C

	//字符串转换
	s2 := strconv.Itoa(10) //数字转换成了字符串
	t.Log(s2)
	//字符串转换为数字
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
