package main

import (
	"fmt"
	"go-dev/algorithm"
)

func main() {
	low, high := 3, 7
	output := algorithm.CountOdds(low, high)
	fmt.Println(output)
}
