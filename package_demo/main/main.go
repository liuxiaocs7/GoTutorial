package main

// 当你的代码都放在 $GOPATH 目录下的话
// 我们导入包要从 $GOPATH/src 后面继续写起
// 给包取别名 heiheihei
// Go语言中不允许导入包而不使用
// Go语言中不允许循环引用包
import (
	"fmt"
	heiheihei "package_demo/calc" // 给导入的包起别名
)

// 多用来做一些初始化的操作
func init() {
	fmt.Println("main.init()")
}

func main() {
	fmt.Println("hello")
	fmt.Println(heiheihei.Name)
	ret := heiheihei.Add(10, 20)
	fmt.Println(ret)
}
