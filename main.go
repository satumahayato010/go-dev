package main

import (
	"fmt"
	"go-dev/algorithm"
)

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	output := algorithm.SortedSquares(nums)
	fmt.Println(output)
}
