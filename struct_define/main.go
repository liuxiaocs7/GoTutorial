package main

import "fmt"

// 定义结构体
// 定义一个person结构体，三个属性
type person struct {
	// 类型相同使用类型简写
	name, city string
	age        int8
}

func main() {
	var p person
	p.name = "沙河娜扎"
	p.city = "北京"
	p.age = 18
	fmt.Printf("%T\n", p)
	// p = main.person{name:"沙河娜扎", city:"北京", age:18}
	// main包里面的一个person类型
	fmt.Printf("p = %#v\n", p)
	// 访问结构体中的内容
	fmt.Println(p.name)
	fmt.Println(p.city)
	fmt.Println(p.age)

	// 匿名结构体
	var user struct {
		name    string
		married bool
	}
	user.name = "小王子"
	user.married = false
	// struct { name string; married bool }{name:"小王子", married:false}
	fmt.Printf("%#v\n", user)
}
