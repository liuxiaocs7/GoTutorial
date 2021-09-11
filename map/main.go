package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

// map(映射)
func main() {
	// 声明一个key为string，值为int的map类型的变量a
	// 光声明map类型，但是没有初始化，a就是初始值nil
	var a map[string]int
	fmt.Println(a == nil) // true

	// map的初始化
	a = make(map[string]int, 8)
	fmt.Println(a == nil) // false

	// map中添加键值对
	a["沙河娜扎"] = 100
	a["沙河小王子"] = 200
	fmt.Println(a) // map[沙河娜扎:100 沙河小王子:200]
	// 同时打印类型和值信息
	fmt.Printf("a:%#v\n", a)   // a:map[string]int{"沙河娜扎":100, "沙河小王子":200}
	fmt.Printf("type:%T\n", a) // type:map[string]int
	// 声明map的同时完成初始化
	b := map[int]bool{
		1: true,
		2: false,
	}
	fmt.Println(b)
	fmt.Printf("b:%#v\n", b) // b:map[int]bool{1:true, 2:false}
	fmt.Printf("type:%T", b) // type:map[int]bool
	fmt.Println()

	// var c map[int]int
	// c[100] = 200 // c则会个map没有初始化，不能直接操作
	// fmt.Println(c)

	// 判断某个键存不存在
	var scoreMap = make(map[string]int, 8)
	scoreMap["沙河娜扎"] = 100
	scoreMap["沙河小王子"] = 200

	// 判断 张二狗子 在不在 scoreMap 中
	// 如果map中存在该键，则value是对应的值
	// 如果map中不存在该键，则value是值类型的零值，同时ok返回一个false
	// value, ok := scoreMap["张二狗子"]
	value, ok := scoreMap["沙河娜扎"]
	fmt.Println(value, ok)
	if ok {
		fmt.Println("张二狗子在scoreMap中", value)
	} else {
		fmt.Println("查无此人")
	}

	// map的遍历
	// for range
	// map是无序的，键值对和添加的顺序无关
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	// 只遍历map中的key
	for k := range scoreMap {
		fmt.Println(k)
	}

	// 只遍历map中的value
	for _, v := range scoreMap {
		fmt.Println(v)
	}

	// 删除 沙河小王子这个键值对
	delete(scoreMap, "沙河小王子")
	fmt.Println(scoreMap)

	// 按照某个固定顺序遍历map
	var scoreMap1 = make(map[string]int, 100)
	// 添加50个键值对
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100) // 0-99的随机整数
		scoreMap1[key] = value
	}
	for k, v := range scoreMap1 {
		fmt.Println(k, v)
	}

	// 按照key从小到大的顺序去遍历scoreMap
	// 1. 先取出所有的key存放到切片中
	// 注意这里不能写为 keys := make([]string, 100)
	// 这样写表示切片中有100个空串，这里的100是len，那么cap也会被设置为100
	// 必须指定len = 0, cap = 100，然后再往里面append元素
	keys := make([]string, 0, 100)
	for k := range scoreMap1 {
		keys = append(keys, k)
	}

	// 2. 对key做排序
	sort.Strings(keys)

	// 3. 按照排序后的key对scoreMap1排序
	for _, key := range keys {
		fmt.Println(key, scoreMap1[key])
	}

	fmt.Println("========================================")
	// 另一种写法尝试，存在问题，大小不能提前固定，只能往后增加，但是结果为什么会在后50个呢？
	keys1 := make([]string, 100, 100)
	cnt := 0
	for k := range scoreMap1 {
		keys1[cnt] = k
		cnt++
		fmt.Println(cnt, k)
	}
	sort.Strings(keys1)
	for num, key := range keys1 {
		fmt.Println(num, key, scoreMap1[key])
	}
	fmt.Println("========================================")

	// 元素类型为map的切片
	// mapSlice 是一个map，但是其类型为map
	var mapSlice = make([]map[string]int, 8, 8) // 只是完成了切片的初始化，内部的map没有进行初始化
	// [nil, nil, nil, nil, nil, nil, nil, nil]
	fmt.Println(mapSlice == nil) // false
	// 还需要完成内部map元素的初始化
	fmt.Println(mapSlice[0] == nil) // true
	// mapSlice[0]["沙河小王子"] = 100
	// fmt.Println(mapSlice)

	// 内层的map必须进行初始化
	mapSlice[0] = make(map[string]int, 8)
	mapSlice[0]["沙河小王子"] = 100
	fmt.Println(mapSlice)

	// 值为切片的map
	// 只完成了map的初始化
	var sliceMap = make(map[string][]int, 8)
	v, ok := sliceMap["中国"]
	if ok {
		fmt.Println(v)
	} else {
		// sliceMap中没有中国这个键
		sliceMap["中国"] = make([]int, 8) // 完成了对切片的初始化
		fmt.Println(sliceMap)
		sliceMap["中国"][0] = 100
		sliceMap["中国"][1] = 200
		sliceMap["中国"][2] = 300
	}

	// 遍历sliceMap
	for k, v := range sliceMap {
		fmt.Println(k, v)
	}

	// 统计一个字符串中每个单词出现的次数
	// "how do you do"中每个单词出现的次数
	// 0.定义一个map[string]int
	var s = "how do you do"
	var wordCount = make(map[string]int, 10)

	// 1.字符串中都有哪些单词
	words := strings.Split(s, " ")

	// 2.遍历单词做统计
	// 这里使用的就是一个默认值，默认为1的
	// for _, word := range words {
	// 	wordCount[word] += 1
	// }
	for _, word := range words {
		v, ok := wordCount[word]
		if ok {
			// map中有这个单词的统计记录
			wordCount[word] = v + 1
		} else {
			wordCount[word] = 1
		}
	}
	fmt.Println(wordCount)
	for k, v := range wordCount {
		fmt.Println(k, v)
	}
}
