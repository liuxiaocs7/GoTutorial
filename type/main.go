package main

import "fmt"

// 自定义类型和类型别名示例

// 1. 自定义类型
// 基于内置的int类型定义了一个全新的MyInt类型
// MyInt类型具有和int类型相同的一些特点
type MyInt int

// 2. 类型别名
// NewInt 是int类型的别名
type NewInt = int

func main() {
	var i MyInt
	// type:main.MyInt value:0  main是包的名字
	fmt.Printf("type:%T value:%v\n", i, i)

	var x NewInt
	// type:int value:0
	fmt.Printf("type:%T value:%v\n", x, x)
}
