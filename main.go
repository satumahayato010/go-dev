package main

import (
	"fmt"
	"go-dev/algorithm"
)

func main() {
	salary := []int{1000, 2000, 3000}
	output := algorithm.Average(salary)
	fmt.Println(output)
}
