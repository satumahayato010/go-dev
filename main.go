package main

import (
	"fmt"
	"go-dev/algorithm"
)

func main() {
	nums := []int{1, 3, 5, 6}
	target := 7
	output := algorithm.SearchInsert(nums, target)
	fmt.Println(output)

}
