package main

import (
	"fmt"
	"go-dev/code75"
)

func main() {
	nums := []int{1, 2, 3, 4}
	output := code75.RunningSum(nums)
	fmt.Println(output)
}
