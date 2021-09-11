package main

import "fmt"

// 函数
// 定义一个不需要参数也没有返回值的函数：sayHello
func sayHello() {
	fmt.Println("Hello 沙河小王子!")
}

// 定义一个接收string类型的name参数
func sayHello2(name string) {
	fmt.Println("Hello ", name)
}

// 定义接收多个参数的函数并且有一个返回值
func intSum(a int, b int) int {
	ret := a + b
	return ret
}

// 返回值简写方式
// 给返回值命名为 ret，在函数中可以直接使用ret
// Go 语言中支持给返回值命名
func intSum2(a int, b int) (ret int) {
	ret = a + b
	return
}

// 函数接收可变参数，函数可以接收0-多个int类型的参数
// 在参数名后面加 ... 表示可变参数，可变参数在函数体中是切片类型
func intSum3(a ...int) int {
	fmt.Println(a)        // [10 20]
	fmt.Printf("%T\n", a) // []int

	ret := 0
	for _, arg := range a {
		ret = ret + arg
	}
	return ret
}

// 固定参数和可变参数同时出现时，可变参数要放在最后
// go语言的函数中没有默认参数
func intSum4(a int, b ...int) int {
	ret := a
	for _, arg := range b {
		ret = ret + arg
	}
	return ret
}

// Go语言中函数参数类型的简写
func intSum5(a, b int) (ret int) {
	ret = a + b
	return
}

// 定义具有多个返回值的函数
func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

func main() {
	// 函数调用
	sayHello()
	name := "沙河娜扎"
	sayHello2(name)
	sayHello2("沙河小王子")

	r := intSum(10, 20)
	fmt.Println(r)

	// 函数调用
	intSum3(10, 20)

	r1 := intSum3()
	r2 := intSum3(10)
	r3 := intSum3(10, 20)
	r4 := intSum3(10, 20, 30)
	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	fmt.Println(r4)

	r11 := intSum4(0)          // a = 0, b = []
	r12 := intSum4(10)         // a = 10, b = []
	r13 := intSum4(10, 20)     // a = 10, b = [20]
	r14 := intSum4(10, 20, 30) // a = 10, b = [20, 30]
	fmt.Println(r11)
	fmt.Println(r12)
	fmt.Println(r13)
	fmt.Println(r14)

	// 函数调用
	x, y := calc(100, 200)
	fmt.Println(x, y)
}
