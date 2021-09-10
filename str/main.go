package main

import (
	"fmt"
	"strings"
)

// 字符串常见操作
func main() {
	// 1.求字符串长度
	s := "hello"
	fmt.Println(len(s)) // 5
	s2 := "hello沙河"
	fmt.Println(len(s2)) // 11

	// 2.拼接字符串
	fmt.Println(s + s2)
	// 第一个位置使用s替换，第二个位置使用s2替换
	s3 := fmt.Sprintf("%s - %s", s, s2)
	fmt.Println(s3)

	// 3.字符串的分割
	s4 := "how do you do"
	// 使用空格来分割s4这个字符串
	fmt.Println(strings.Split(s4, " "))
	// 输出切割后的类型
	fmt.Printf("%T\n", strings.Split(s4, " "))

	// 4.判断是否包含
	// s4这个字符串是否包含 "do" 这个子串
	fmt.Println(strings.Contains(s4, "do"))

	// 5.判断前缀
	// 判断字符串s4的前缀是不是 how
	fmt.Println(strings.HasPrefix(s4, "how"))

	// 6.判断后缀
	fmt.Println(strings.HasSuffix(s4, "how"))
	fmt.Println(strings.HasSuffix(s4, "do"))

	// 7.判断子串的位置
	fmt.Println(strings.Index(s4, "do"))

	// 8.最后子串出现的位置
	fmt.Println(strings.LastIndex(s4, "do"))

	// 9.join  将字符串切片使用连字符连接起来
	s5 := []string{"how", "do", "you", "do"}
	fmt.Println(s5)
	fmt.Println(strings.Join(s5, "+"))

}
