package main

import "fmt"

// 字符
func main() {
	// byte uint8的别名 ASCII码
	// rune int32的别名
	var c1 byte = 'c'
	var c2 rune = 'c'
	fmt.Println(c1, c2) // 99 99
	// 打印类型
	fmt.Printf("c1:%T  c2:%T", c1, c2) // c1:uint8  c2:int32

	s := "hello沙河"
	// 按照byte取值的
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c\n", s[i])
	}
	fmt.Println()
	// 按照rune取值的
	for _, r := range s {
		fmt.Printf("%c\n", r)
	}
}
