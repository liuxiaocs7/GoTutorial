package main

import (
	"fmt"
	"math"
)

// 基本数据类型

func main() {
	// 十进制打印为二进制
	n := 10
	// %b表示将一个数转为二进制
	fmt.Printf("%b\n", n)
	// %d表示以十进制的形式输出
	fmt.Printf("%d\n", n)

	// 八进制
	m := 075
	fmt.Printf("%d\n", m) // 十进制61
	fmt.Printf("%o\n", m)

	// 十六进制
	f := 0xff
	fmt.Println(f)        // 十进制的255
	fmt.Printf("%x\n", f) // ff

	// uint8
	var age uint8 // 0-255
	// age = 256  溢出直接报错，为对应的数据选择合适的数据类型
	fmt.Println(age)

	// 浮点数
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	var a bool
	fmt.Println(a) // 默认值是 false
	a = true
	fmt.Println(a)

	// 字符串
	s1 := "hello beijing"
	s2 := "你好 北京"
	fmt.Println(s1)
	fmt.Println(s2)

	// 打印windows平台下的一个路径 c:\code\go.exe
	fmt.Println("c:\\code\\go.exe")
	fmt.Println("\t制表符\n换行符")
	s3 := `
	多行字符串
	两个反引号之间的内容
	会
	原样输出
	\t
	\n
	`
	fmt.Println(s3)
}
