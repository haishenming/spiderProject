package main

import "fmt"

func main() {
	// 数组是值类型
	var arr1 [5]int // 定义数组
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{1, 3, 5, 7, 9}

	var grid [4][5]bool // 多维数组

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	// 遍历数组
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
}
