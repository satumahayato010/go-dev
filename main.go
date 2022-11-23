package main

import (
	"fmt"
	"go-dev/data-structure"
)

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	output := data_structure.MaxSubArray(nums)
	fmt.Println(output)
}
