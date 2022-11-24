package main

import (
	"fmt"
	data_structure "go-dev/data-structure"
)

func main() {
	nums1, nums2 := []int{1, 2, 3, 0, 0, 0}, []int{2, 5, 6}
	data_structure.Merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}
