package main

import "fmt"

// 声明全局变量并初始化
var x = 100
var y = "小王子"

// 变量
func main() {
	// 1.标准声明格式
	var name string
	var age int
	var isOk bool

	fmt.Println(name, age, isOk)

	// 2.批量声明变量
	var (
		a string
		b int
		c bool
		d float32
	)
	fmt.Println(a, b, c, d)

	// 3.声明变量同时指定初始值
	var name1 string = "小王子"
	var age1 int = 18
	fmt.Println(name1, age1)
	var name2, age2 = "沙河娜扎", 28
	fmt.Println(name2, age2)

	// 4.类型推导 [可以定义在全局]
	// 编译器根据变量的初始值推导出其类型
	var name3 = "小王子"
	var age3 = 19
	fmt.Println(name3, age3)

	// 5.短变量声明 [使用较多，只能用在函数中]
	// 声明变量并初始化变量
	// 编译器根据这个10推导出m这个变量的类型是int
	m := 10
	fmt.Println(m)

	// 6.匿名变量 使用 _ 接收这个值

}
