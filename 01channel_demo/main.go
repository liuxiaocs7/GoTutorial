package main

import "fmt"

func main() {
	// 引用类型,需要初始化之后才能使用
	// 声明通道ch1
	// var ch1 chan int
	// 初始化ch1,类型为chan int,容量为1
	// ch1 = make(chan int, 1)

	// ch1 := make(chan int)    // 无缓冲区通道,又称为同步通道
	ch1 := make(chan int, 1) // 带缓冲区通道,异步通道
	// 将值存入通道
	ch1 <- 10
	// 从通道中取值
	x := <-ch1
	fmt.Println(x)

	// 取通道中元素的数量
	fmt.Println(len(ch1))
	// 拿到通道的容量
	fmt.Println(cap(ch1))

	// 关闭通道
	close(ch1)
}
