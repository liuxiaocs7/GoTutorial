package main

import (
	"fmt"
	"reflect"
)

// 结构体反射

type student struct {
	Name  string `json:"name" ini:"s_name"`
	Score int    `json:"score" ini:"s_score"`
}

// 给student添加两个方法 Study和Sleep(注意首字母大写)
func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

// 根据反射去获取结构体中方法的函数
func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	// 取到变量的方法数量
	fmt.Println(t.NumMethod())

	v := reflect.ValueOf(x)
	// 因为下面需要拿到具体的方法去调用，所以使用的是值信息
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method: %s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		// 参数必须是一个reflect.Value{}类型的切片
		var args = []reflect.Value{}
		// 调用方法
		v.Method(i).Call(args)
	}

	// 通过方法名获取结构体的方法
	t.MethodByName("Sleep")              // (Method, bool)
	methodObj := v.MethodByName("Sleep") // Value
	fmt.Println(methodObj.Type())
}

func main() {
	stu1 := student{
		Name:  "小王子",
		Score: 90,
	}

	// 通过反射去获取结构体中的所有字段信息
	t := reflect.TypeOf(stu1)
	fmt.Printf("name:%v kind:%v\n", t.Name(), t.Kind())

	// 遍历结构体变量的所有字段
	for i := 0; i < t.NumField(); i++ {
		// 根据结构体字段的索引去取字段
		fieldObj := t.Field(i)
		fmt.Printf("name:%v type:%v tag:%v\n", fieldObj.Name, fieldObj.Type, fieldObj.Tag)
		fmt.Println(fieldObj.Tag.Get("json"), fieldObj.Tag.Get("ini"))
	}

	// 根据名字取结构体中的字段
	fieldObj2, ok := t.FieldByName("Score")
	if ok {
		fmt.Printf("name:%v type:%v tag:%v\n", fieldObj2.Name, fieldObj2.Type, fieldObj2.Tag)
	}

	printMethod(stu1)
}
