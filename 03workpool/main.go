package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, job)
		result <- job * 2
		time.Sleep(time.Microsecond * 500)
		fmt.Printf("worker:%d end job:%d\n", id, job)
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 开启3个goroutine
	for j := 0; j < 3; j++ {
		go worker(j, jobs, results)
	}

	// 发送5个任务
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)
	// 输出结果
	for i := 0; i < 5; i++ {
		ret := <-results
		fmt.Println(ret)
	}
	// results 必须要关掉才能这样取值,否则会死锁
	// for ret := range results {
	// 	fmt.Println(ret)
	// }
	time.Sleep(time.Second)
}
