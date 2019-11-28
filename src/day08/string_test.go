package day08

import (
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestString(t *testing.T) {
	var s string //初始化默认值是""
	t.Log(len(s), s)
	s = "hello"
	t.Log(len(s))
	//s[1] = '2' //string 是不可变的byte slice
	s = "\xE4\xB8\xA5" //存储二进制数据
	t.Log(len(s), s)
	s = "中"
	t.Log(len(s)) //byte数
}

func TestLenStr(t *testing.T) {
	s := "hello 你好"
	t.Log(len(s))         // 12 字符串所占字节的长度
	t.Log(len([]byte(s))) // 12 字符串所占字节的长度
	//计算字符串的长度
	//golang中的unicode/utf8包提供了用utf-8获取长度的方法
	t.Log(utf8.RuneCountInString(s)) //8
	//通过rune类型处理unicode字符
	t.Log(len([]rune(s))) //8
	s = "中"
	c := []rune(s)
	t.Logf("中 的Unicode %x", c[0])
	t.Logf("中 的utf-8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华有为-华为"
	//字符串range遍历自动化获取rune,而不是byte
	for _, c := range s {
		// %[1]c %[1]d代表都是以第一个遍历进行格式化输出
		t.Logf("%[1]c %[1]d", c)
		//t.Logf("%[1]c %[1]x", c)
	}
}

func TestStringFn(t *testing.T) {
	s := "A-B-C"
	// Split分割获得切片
	parts := strings.Split(s, "-")
	for _, part := range parts {
		t.Logf("%s", part)
	}
	// 拼接
	str := strings.Join(parts, ",")
	t.Log(str)
}

func TestStrConv(t *testing.T) {
	// 字符串和int转换
	s := strconv.Itoa(1)
	t.Log("字符串" + s)
	//t.Log(10 + strconv.Atoi("1"))
	//Atoi返回值是两个，不能直接相加
	if n, err := strconv.Atoi("1"); err == nil {
		t.Log(10 + n)
	}
}
