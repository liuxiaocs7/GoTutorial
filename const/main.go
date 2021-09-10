package main

import "fmt"

// 常量
// const pi = 3.1415
// const e = 2.7

// 批量声明
// const (
// 	pi = 3.1415
// 	e  = 2.7
// )

// 如果某一行后面没有跟表达式
// 那么其值和上面一样
// const (
// 	n1 = 10
// 	n2
// 	n3
// )

// iota 自动赋值
// const (
// 	n1 = iota // 0
// 	n2        // 1
// 	_         // 忽略2这个值
// 	n4        // 3
// )

const (
	n1 = iota // 0
	n2 = iota // 1
	n3 = 100  // 100  [使用100覆盖了默认的值2]
	n4 = iota // 3
)

// 遇到const，iota就变成了0
const n5 = iota

const (
	_  = iota
	KB = 1 << (10 * iota) // 1 << 10，1024
	MB = 1 << (10 * iota) // 1 << 20
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

const (
	a, b = iota + 1, iota + 2 // iota = 0: 1,2
	c, d                      // iota = 1: 2,3
	e, f                      // iota = 2: 3,4
)

func main() {
	// fmt.Println(pi, e)
	// 0 1 100 3 0
	fmt.Println(n1, n2, n3, n4, n5)
	// 1 2 2 3 3 4
	fmt.Println(a, b, c, d, e, f)
}
