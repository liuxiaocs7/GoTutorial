package main

import "fmt"

// 数组
func main() {
	// 数组默认初始值为0
	var a [3]int
	var b [4]int
	// a = b  // 不能直接这样赋值，因为两者类型不同
	fmt.Println(a)
	fmt.Println(b)

	// 数组的初始化
	// 1. 定义时使用初始值列表的方式初始化
	// var cityArray = [4]string{"北京", "上海", "广州", "深圳"}
	// fmt.Println(cityArray)
	// fmt.Println(cityArray[0])
	// fmt.Println(cityArray[3])

	// 2. 编译器推导数组的长度
	// var boolArray = [...]bool{true, false, false, true, false}
	// fmt.Println(boolArray)
	// fmt.Printf("%T\n", boolArray) // [5]bool

	// 3. 使用索引值方式初始化
	// 指定三个索引值
	// var langArray = [...]string{1: "Golang", 3: "Python", 7: "Java"}
	// fmt.Println(langArray)        // [ Golang  Python    Java]
	// fmt.Printf("%T\n", langArray) // [8]string

	// 数组的遍历
	var cityArray = [4]string{"北京", "上海", "广州", "深圳"}
	// 1. for 循环遍历
	for i := 0; i < len(cityArray); i++ {
		fmt.Println(cityArray[i])
	}

	// 2. for range 遍历
	// 2.1 同时获取索引和值
	for index, value := range cityArray {
		fmt.Println(index, value)
	}

	// 2.2 只获取索引
	for index := range cityArray {
		fmt.Println(index)
	}

	// 2.3 同时获取索引和值
	for _, value := range cityArray {
		fmt.Println(value)
	}

	// 二维数组
	// 多维数组中只有最外层可以使用多个点的简写方式让编译器推导长度
	cityArray1 := [4][2]string{
		{"北京", "西安"},
		{"上海", "杭州"},
		{"重庆", "成都"},
		{"广州", "深圳"},
	}
	fmt.Println(cityArray1)
	fmt.Println(cityArray1[2][0])

	fmt.Println()

	// 二维数组的遍历
	for _, v1 := range cityArray1 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	// 数组是值类型
	x := [3]int{1, 2, 3}
	fmt.Println(x) // [1 2 3]
	f1(x)
	fmt.Println(x) // [1 2 3]

	y := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(y)
	f2(y)
	fmt.Println(y)

	// z相当于是对y的值拷贝
	// 完完整整拷贝一份赋给新的变量
	z := y
	z[0][0] = 1000
	fmt.Println(y)
	fmt.Println(z)
}

func f1(a [3]int) {
	a[0] = 100
}

func f2(a [3][2]int) {
	a[0][0] = 100
}
