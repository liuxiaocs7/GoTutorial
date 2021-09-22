package main

import "fmt"

// 如果又多个条件满足, select会随机选择一种执行
func main() {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		default:
			fmt.Println("啥都不干")
		}
	}
}
