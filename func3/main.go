package main

import (
	"fmt"
	"strings"
)

// 定义一个函数它的返回值是一个函数
// 把函数作为返回值
func a() func() {
	name := "沙河娜扎"
	// 先在函数内部找，然后再在外层函数中找
	return func() {
		fmt.Println("hello, 沙河小王子")
		fmt.Println("hello", name)
	}
}

func a1(name string) func() {
	// 先在函数内部找，然后再在外层函数中找
	return func() {
		fmt.Println("hello", name)
	}
}

// 返回的类型是一个函数，这个函数的参数是string类型的，这个函数的返回值是string类型的
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

// 匿名函数和闭包
func main() {
	sayHello := func() {
		fmt.Println("匿名函数")
	}
	// 调用匿名函数
	sayHello()

	// 定义并执行匿名函数
	func() {
		fmt.Println("匿名函数")
	}()

	// 获取函数
	// 闭包 = 函数 + 外层变量的引用
	r := a() // r此时就是一个闭包
	// 执行函数，相当于执行了a函数内部的匿名函数
	r()

	r1 := a1("沙河小王子")
	r1()

	r2 := makeSuffixFunc(".txt")
	ret := r2("沙河娜扎")
	fmt.Println(ret)

	r3 := makeSuffixFunc(".avi")
	ret3 := r3("沙河娜扎")
	fmt.Println(ret3)

	x, y := calc(100)
	retx := x(200) // base = 100 + 200
	fmt.Println(retx)
	rety := y(200) // base = 300 - 200
	fmt.Println(rety)
}
