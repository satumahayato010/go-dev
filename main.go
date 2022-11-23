package main

import (
	"fmt"
	data_structure "go-dev/data-structure"
)

func main() {
	nums := []int{3, 3}
	target := 6
	output := data_structure.TwoSum(nums, target)
	fmt.Println(output)
}
