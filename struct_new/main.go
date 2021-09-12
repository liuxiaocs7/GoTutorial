package main

import "fmt"

// 结构体构造函数：构造一个结构体实例的函数
type person struct {
	name, city string
	age        int8
}

// 创建结构体
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

func main() {
	p1 := newPerson("小王子", "北京", int8(18))
	// type:*main.person value:&main.person{name:"小王子", city:"北京", age:18}
	fmt.Printf("type:%T value:%#v\n", p1, p1)
}
