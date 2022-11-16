package main

import "fmt"

func twoSum(nums []int, target int) []int {
	lookup := make(map[int]int)

	for idx, num := range nums {
		if pos, ok := lookup[target-num]; ok {
			return []int{pos, idx}
		}
		lookup[num] = idx
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	output := twoSum(nums, target)
	fmt.Println(output)
}
