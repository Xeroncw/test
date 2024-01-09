package main

import (
	"fmt"
)

func main() {
	var arr1 [5] int
	fmt.Println(arr1)

	var arr2 = [5] int {1, 2, 3, 4, 5}
	fmt.Println(arr2)

	var arr3 = [...] int {1, 2, 3, 4, 5, 6}
	fmt.Println(arr3)

	arr4 := [...] int {1, 2, 3, 4, 5}
	fmt.Println(arr4)

	arr5 := [5] int {1, 2, 3, 4, 5}
	fmt.Println(arr5)

	//二维数组
	var arr6 = [3][5] int {{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}}
	fmt.Println(arr6)

	arr7 := [3][5] int {{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}}
	fmt.Println(arr7)

	arr8 := [...][5] int {{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}}
	fmt.Println(arr8)
}