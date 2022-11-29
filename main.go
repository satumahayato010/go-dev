package main

import (
	"fmt"
	"go-dev/algorithm"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	algorithm.Rotate(nums, k)
	fmt.Println(nums)
}
