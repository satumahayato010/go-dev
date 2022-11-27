package main

import (
	"fmt"
	data_structure "go-dev/data-structure"
)

func main() {
	nums1, nums2 := []int{4, 9, 5}, []int{9, 4, 9, 8, 5}
	output := data_structure.Intersect(nums1, nums2)
	fmt.Println(output)
}
