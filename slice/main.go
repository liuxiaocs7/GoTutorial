package main

import (
	"fmt"
	"sort"
)

// 切片(slice)
// 切片要初始化之后才能使用
func main() {
	// 定义一个字符串类型的切片
	// var a []string
	// 定义一个int类型的切片
	// var b []int
	// 定义一个bool类型的切片并初始化
	// var c = []bool{false, true}
	// fmt.Println(a) // []
	// fmt.Println(b) // []
	// fmt.Println(c) // [false true]

	// 基于数组得到切片
	a := [5]int{55, 56, 57, 58, 59}
	b := a[1:4]
	fmt.Println(b)        // [56 57 58]
	fmt.Printf("%T\n", b) // []int

	// 切片再次切片
	// b[0:] 等价于 b[0:len(b)]
	c := b[0:]
	fmt.Println(c)
	fmt.Printf("%T\n", c) // []int

	// make函数构造切片
	// 第一个参数是切片的类型
	// 第二个参数是切片的长度
	// 第三个参数是切片底层数组的长度(即容量)，如果省略则和长度一致
	// d是一个长度为5，容量为10的一个int类型的切片
	d := make([]int, 5, 10)
	fmt.Println(d)        // [0 0 0 0 0]
	fmt.Printf("%T\n", d) // []int

	// 通过len()函数获取切片的长度
	fmt.Println(len(d)) // 5
	// 通过cap()函数获取切片的容量
	fmt.Println(cap(d)) // 10

	// nil
	var a1 []int         // 声明int类型切片
	var b1 = []int{}     // 声明并初始化，这里的b1不是一个nil，已经在底层创建了一个数组与之对应
	c1 := make([]int, 0) // 同样地，这里的c1也不是nil
	if a1 == nil {
		fmt.Println("a == nil")
	}
	fmt.Println(a1, len(a1), cap(a1)) // [] 0 0

	if b == nil {
		fmt.Println("b == nil")
	}
	fmt.Println(b1, len(b1), cap(b1)) // [] 0 0

	if c == nil {
		fmt.Println("c == nil")
	}
	fmt.Println(c1, len(c1), cap(c1)) // [] 0 0

	// 切片的赋值拷贝
	// a2, b2 共同指向了内存中的同一个数组
	a2 := make([]int, 3) // [0, 0, 0]
	b2 := a2
	b2[0] = 100
	fmt.Println(a2) // [100 0 0]
	fmt.Println(b2) // [100 0 0]

	// 切片的遍历
	a3 := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(a3); i++ {
		fmt.Println(i, a3[i])
	}
	fmt.Println()

	for index, value := range a3 {
		fmt.Println(index, value)
	}

	// 切片的扩容
	var a4 []int // 此时它并没有申请内存，只是相当于声明了一个变量
	// 因为有可能返回一个新的地址，因此需要使用a4来接收append()函数的返回值，涉及到底层的扩容
	// a4 = append(a4, 10)
	// fmt.Println(a4)
	// for i := 0; i < 10; i++ {
	// 	a4 = append(a4, i)
	// 	// %v值 %d整数 %p指针地址
	// 	fmt.Printf("%v len:%d cap:%d ptr:%p\n", a4, len(a4), cap(a4), a4)
	// }
	// 一次追加多个元素
	a4 = append(a4, 1, 2, 3, 4, 5)
	fmt.Println(a4)
	b4 := []int{12, 13, 14, 15}
	// b4... 表示逐个取出切片中的元素
	a4 = append(a4, b4...)
	fmt.Println(a4)

	// 切片的copy
	a5 := []int{1, 2, 3, 4, 5}
	b5 := make([]int, 5, 5)
	c5 := b5
	// a5和b5不是指向同一个位置了
	copy(b5, a5)
	b5[0] = 100
	fmt.Println(a5)
	fmt.Println(b5)
	fmt.Println(c5)

	// 切片删除元素
	a6 := []string{"北京", "上海", "广州", "深圳"}
	// append(a[:index], a[index+1:...])
	a6 = append(a6[0:2], a6[3:]...)
	fmt.Println(a6)

	// 练习
	var a7 = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a7 = append(a7, fmt.Sprintf("%v", i))
	}
	// len = 5，前面5个是空串
	// 然后往后面添加10个元素的时候先将cap占满，之后再扩容
	fmt.Println(a7, len(a7), cap(a7)) // [     0 1 2 3 4 5 6 7 8 9] 15 20

	// 请使用内置的sort包队数组var a = [...]int{3, 7, 8, 9, 1}进行排序。
	var a8 = [...]int{3, 7, 8, 9, 1}
	// 因为sort排序只接受一个切片作为参数，因此这里需要将数组转为切片
	sort.Ints(a8[:])
	fmt.Println(a8)
}
