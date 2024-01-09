package main

import (
	"fmt"
)

func main() {
	var sli1 [] int
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli1),cap(sli1),sli1)

	var sli2 = [] int {}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli2),cap(sli2),sli2)

	var sli3 = [] int {1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli3),cap(sli3),sli3)

	sli4 := [] int {1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli4),cap(sli4),sli4)

	var sli5 = make([] int, 5, 8)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli5),cap(sli5),sli5)

	sli6 := make([] int, 5, 9)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli6),cap(sli6),sli6)
}