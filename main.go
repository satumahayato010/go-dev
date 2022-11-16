package main

import "fmt"

func totalNumericArray(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}

func arrayMaximumValue(nums []int) int {
	max := nums[0]
	for i, _ := range nums {
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}

func main() {
	nums := []int{7, 3, 5, 1, 8, 4}
	output := totalNumericArray(nums)
	fmt.Println(output)
	result := arrayMaximumValue(nums)
	fmt.Println(result)
}
