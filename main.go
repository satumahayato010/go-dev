package main

import (
	"fmt"
	"math/rand"
)

func inOrder(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}

func shuffle(nums []int) {
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
}

func bogoSort(nums []int) {
	for !inOrder(nums) {
		shuffle(nums)
	}
}

func main() {
	nums := []int{1, 5, 3, 2, 6}
	bogoSort(nums)
	fmt.Println(nums)
}
