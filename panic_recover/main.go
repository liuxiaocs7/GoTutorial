package main

import "fmt"

func a() {
	fmt.Println("func a")
}

func b() {
	// 在函数b执行完之前执行
	// recover()函数从当前程序运行环境中将错误信息搜集起来
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("func b error")
		}
	}()
	panic("panic in b")
}

func c() {
	fmt.Println("func c")
}

// panic和recover
func main() {
	a()
	b()
	c()
}
