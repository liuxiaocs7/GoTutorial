package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("hello", i)
	// 向全局的wg通知已经干完活了，通过wg把计数器-1
	wg.Done()
}

// 开启一个主goroutine去执行main函数
func main() {

	// 计数牌+1: 表示开启一个子goroutine
	// wg.Add(1)
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		// wg.Add(1)
		// 开启了一个goroutine去执行hello这个函数
		go hello(i)
	}

	fmt.Println("hello main")

	// 让所有的goroutine都执行完
	// time.Sleep(time.Second)
	// 阻塞，等所有小弟都干完活之后才收兵，等到计数器==0的时候才退出
	wg.Wait()
}
