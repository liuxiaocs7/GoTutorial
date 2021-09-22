package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// 开启一个主goroutine去执行main函数
func main() {

	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		// 闭包,变量i引用外部的变量
		// 解决方法是将变量i作为参数传入进去
		go func(i int) {
			fmt.Println("hello", i)
			wg.Done()
		}(i)
	}

	fmt.Println("hello main")
	wg.Wait()
}
