package main

import "fmt"

// 结构体的初始化
type person struct {
	name, city string
	age        int8
}

func main() {
	// 1. 键值对初始化
	p4 := person{
		name: "小王子",
		age:  18,
	}
	// main.person{name:"小王子", city:"", age:18}
	fmt.Printf("%#v\n", p4)

	p5 := &person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	// &main.person{name:"小王子", city:"北京", age:18}
	fmt.Printf("%#v\n", p5)

	// 2. 值的列表进行初始化
	// 一定要和定义结构体中的顺序一一对应
	p6 := person{
		"小王子",
		"北京",
		18, // 后面的这个逗号一定不能忘了
	}
	fmt.Printf("%#v\n", p6)
}
