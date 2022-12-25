package main

import (
	"fmt"
	"go-dev/algorithm"
)

func main() {
	nums := []int{2, 3, 4}
	target := 9
	output := algorithm.TwoSumDifficult(nums, target)
	fmt.Println(output)
}
