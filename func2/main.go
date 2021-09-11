package main

import "fmt"

// 函数进阶之变量作用域

// 定义全局变量num
var num = 10

// 定义函数
func testGlobal() {
	// 优先使用函数内部的同名变量
	// 可以在函数中使用变量
	// 1. 先在自己函数中查找，找到了就用自己函数中的
	// 2. 函数中找不到变量就往外找全局变量
	num := 100
	name := "沙河娜扎"
	// 可以在函数中访问全局变量
	fmt.Println("变量num", num)
	fmt.Println(name)
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

// op是一个参数，函数类型的，两个int作为参数，返回值为int
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func main() {
	testGlobal()
	// 外层不能访问到函数的内部变量(局部变量)
	// fmt.Println(name)
	fmt.Println("全局变量", num)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 变量i此时只在for循环的语句块中生效
	// fmt.Println(i) // 外层访问不到内部for语句块中的变量

	// 将函数作为变量直接赋值
	// 函数是可以作为变量
	abc := testGlobal
	fmt.Printf("%T\n", abc)
	abc()

	r1 := calc(100, 200, add)
	fmt.Println(r1)
	r2 := calc(100, 200, sub)
	fmt.Println(r2)
}
