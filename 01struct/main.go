package main

import "fmt"

// 结构体的匿名字段

type Person struct {
	string
	int
}

func main() {
	p1 := Person{
		"小王子",
		18,
	}
	fmt.Println(p1)
	// 因为字段没有name这个属性，因此需要根据类型来访问
	fmt.Println(p1.string, p1.int)
}
