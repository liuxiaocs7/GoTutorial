// 非注释的第一行为包的声明，定义一个包的名字
// main包是一个可以独立执行的程序，每个Go应用程序都需要包含一个main包
package main

// 需要 fmt 这个包
import "fmt"

// 定义一个函数
// 花括号不能单独放一行
// 一行代码就表示结束，末尾不需要添加分号
func main() {
	// 这是单行注释
	/*
		多行注释/块注释
	*/
	// 调用 fmt 包里面的 Println 函数
	fmt.Println("hello 沙河娜扎！")
}
