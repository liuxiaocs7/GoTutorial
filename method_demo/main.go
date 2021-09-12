package main

import "fmt"

// 方法的定义实例
// Person 是一个结构体，并且首字母大写，对外部是可见的
type Person struct {
	name string
	age  int8
}

// 构造函数，通常是用new开头的，是一个Person类型的构造函数
// 函数是不属于任何类型的，是和类型无关的，想在哪里调用就在哪里调用
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// 定义方法
// 函数的接收者：我这个函数是给谁定义的，什么类型的变量会有一个这样的方法
// Person表示接收者的类型，p是接收者，通常是类型的首字母小写
// p是具体的某个接收者类型的实例
// Dream 是为Person类型定义方法
// 但是方法是属于某个类型的，p1 作为接收者参数传入方法的形参
// 使用值接收者相当于拷贝一份p传进来
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

// SetAge 是一个修改年龄的方法
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

// SetAge2 是一个使用值接收者来修改年龄的方法
// 这种写法相比于上一种写法就不会修改
func (p Person) SetAge2(newAge int8) {
	p.age = newAge
	fmt.Println("Inner age: ", p.age, newAge)
}

func main() {
	p1 := NewPerson("沙河娜扎", int8(18))
	// (*p1) 获取对应的对象，然后通过对象调用方法
	(*p1).Dream()
	// 因为不会涉及到指针的操作，所以可以直接写 p1.Dream()
	p1.Dream()

	// 调用修改年龄的方法
	fmt.Println(p1.age) // 18
	p1.SetAge(int8(20))
	fmt.Println(p1.age) // 20

	p1.SetAge2(int8(30))
	fmt.Println(p1.age) // 20
}
