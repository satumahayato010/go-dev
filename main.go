package main

import (
	"fmt"
	"go-dev/data-structure"
)

func main() {
	nums := []int{1, 2, 3, 3}
	output := data_structure.ContainsDuplicate(nums)
	fmt.Println(output)
}
