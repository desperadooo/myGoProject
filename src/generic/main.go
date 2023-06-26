package main

import (
	"fmt"
)

func print[T any] (arr []T) {
	for _, v := range arr {
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Println("")
}

func main() {
	strs := []string{"A", "B", "C"}
	decs := []float64{3.14, 2.56, 8.91}
	nums := []int{1, 2, 3}
	print(strs)
	print(decs)
	print(nums)
}