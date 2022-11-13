package main

import "fmt"

func twoSum(nums []int, target int) []int {
	lookup := make(map[int]int)

	for i, v := range nums {
		j, ok := lookup[-v]
		lookup[v-target] = i

		if ok {
			return []int{j, i}
		}
	}
	return []int{}
}

func main() {
	//nums := []int{2, 7, 11, 15}
	nums2 := []int{3, 2, 4}
	target := 6
	output := twoSum(nums2, target)
	fmt.Println(output)
}
