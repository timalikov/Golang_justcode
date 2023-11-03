package main

import (
	"Lecture01/task1"
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := task1.TwoSum(nums, target)
	fmt.Println(result)

}
