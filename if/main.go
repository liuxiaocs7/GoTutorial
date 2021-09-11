package main

import "fmt"

// if判断
func main() {
	// 1.基本写法
	// var score = 65
	// if score >= 90 {
	// 	fmt.Println("A")
	// } else if score > 75 {
	// 	fmt.Println("B")
	// } else {
	// 	fmt.Println("C")
	// }

	// 2.if判断的特殊写法
	// score这个变量只会在if这个块内生效
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
	// fmt.Println(score)  // if外面无法访问score

	score := 100
	if score > 95 {
		fmt.Println("A+")
	}
}
