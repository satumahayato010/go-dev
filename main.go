package main

import (
	"fmt"
	"go-dev/algorithm"
)

func main() {
	nums1 := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	low, high := 0, len(nums1)-1
	output := algorithm.BinarySearchRecursive(nums1, target, low, high)
	fmt.Println(output)
}
