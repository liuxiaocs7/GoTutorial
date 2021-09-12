package main

import "fmt"

// 结构体指针

type person struct {
	name, city string
	age        int8
}

func main() {
	// 返回一个指针
	var p2 = new(person)
	// *main.person
	fmt.Printf("%T\n", p2)

	// 赋值
	// (*p2).name = "小王子"
	// (*p2).city = "北京"
	// (*p2).age = 18
	// 按照上面的方法写看起来比较复杂，可以简写为 指针名.字段名 的形式
	p2.name = "小王子"
	p2.city = "北京"
	p2.age = 18
	// &main.person{name:"小王子", city:"", age:18}
	fmt.Printf("%#v\n", p2)

	// 取结构体的地址进行实例化
	// person{} 表示实例化了一个person的结构体
	p3 := &person{}
	// main.person
	fmt.Printf("%T\n", p3)
	// main.person{name:"", city:"", age:0}
	// 可见有默认初始化
	fmt.Printf("%#v\n", p3)
	p3.name = "娜扎"
	p3.city = "北京"
	p3.age = 18
	// &main.person{name:"娜扎", city:"北京", age:18}
	fmt.Printf("%#v", p3)
}
