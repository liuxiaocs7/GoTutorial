package main

import "fmt"

func main() {
	var a *int      // a = nil
	a = new(int)    // 初始化
	fmt.Println(*a) // 0
	*a = 100
	fmt.Println(*a) // 100

	var b map[string]int     // map = nil
	b = make(map[string]int) // 初始化
	b["沙河娜扎"] = 100
	fmt.Println(b)
}
