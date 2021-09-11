package main

import "fmt"

func main() {
	// 1.求数组
	a := [...]int{1, 3, 5, 7, 9}
	sumValue := 0
	for _, value := range a {
		sumValue += value
	}
	fmt.Println(sumValue)

	// 2. 找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0, 3)和(1,2)
	b := [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(b); i++ {
		for j := i + 1; j < len(b); j++ {
			if b[i]+b[j] == 8 {
				fmt.Println(i, j)
			}
		}
	}
}
