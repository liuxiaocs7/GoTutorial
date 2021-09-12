package main

import (
	"encoding/json"
	"fmt"
)

// 结构体字段的可见性与JSON序列化

// 如果一个GO语言包中定义的标识符是首字母大写的，那么就是对外可见的
// 如果是小写的则只能在本包内使用
// 如果一个结构体中的字段名首字母是大写的，那么它该字段就是对外可见的
type student struct {
	ID   int
	Name string
}

// student的构造函数
func newStudent(id int, name string) student {
	return student{
		ID:   id,
		Name: name,
	}
}

// 如果这个类型下面的属性不是大写的，那么json包中的方法是无法访问title和students的
// `json:"title"` 改变序列化之后字符串中的命名，严格按照这个格式！！！
// db:"student" 数据库查询的时候使用小写的 student
type class struct {
	Title    string    `json:"title"`
	Students []student `json:"student_list" db:"student" xml:"ss"`
}

func main() {
	// 创建一个班级变量c1
	c1 := class{
		Title:    "火箭101",
		Students: make([]student, 0, 20), // 切片初始化
	}
	// 往班级c1中添加学生
	for i := 0; i < 10; i++ {
		// 造十个学生
		tmpStu := newStudent(i, fmt.Sprintf("stu%02d", i))
		c1.Students = append(c1.Students, tmpStu)
	}

	// main.class{Title:"火箭101", Students:[]main.student{
	// main.student{ID:0, Name:"stu00"}, main.student{ID:1, Name:"stu01"},
	// main.student{ID:2, Name:"stu02"}, main.student{ID:3, Name:"stu03"},
	// main.student{ID:4, Name:"stu04"}, main.student{ID:5, Name:"stu05"},
	// main.student{ID:6, Name:"stu06"}, main.student{ID:7, Name:"stu07"},
	// main.student{ID:8, Name:"stu08"}, main.student{ID:9, Name:"stu09"}}}
	fmt.Printf("%#v\n", c1)

	// JSON序列化： Go语言中的数据 -> JSON格式的字符串
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json marshal failed, err: ", err)
		return
	}
	fmt.Printf("%T\n", data) // []uint8
	fmt.Printf("%s\n", data) // {"Title":"火箭101","Students":[{"ID":0,"Name":"stu00"},{"ID":1,"Name":"stu01"},{"ID":2,"Name":"stu02"},{"ID":3,"Name":"stu03"},{"ID":4,"Name":"stu04"},{"ID":5,"Name":"stu05"},{"ID":6,"Name":"stu06"},{"ID":7,"Name":"stu07"},{"ID":8,"Name":"stu08"},{"ID":9,"Name":"stu09"}]}

	// JSON反序列化：JSON格式的字符串 -> Go语言中的数据
	jsonStr := `{"Title":"火箭101","Students":[{"ID":0,"Name":"小王子"},{"ID":1,"Name":"哪吒"},{"ID":2,"Name":"stu02"}]}`

	// 只有传指针才能修改内部的值
	var c2 class
	// 将字符串强转为 []byte 格式，并且将转换之后的数据存储到c2中
	err = json.Unmarshal([]byte(jsonStr), &c2)
	if err != nil {
		fmt.Println("json unmarshal")
		return
	}
	// main.class{Title:"火箭101", Students:[]main.student{main.student{ID:0, Name:"stu00"}, main.student{ID:1, Name:"stu01"}, main.student{ID:2, Name:"stu02"}}}
	fmt.Printf("%#v\n", c2)
}
