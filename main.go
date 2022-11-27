package main

import (
	"fmt"
	data_structure "go-dev/data-structure"
)

func main() {
	prices := []int{7, 6, 4, 3, 1}
	output := data_structure.MaxProfit(prices)
	fmt.Println(output)
}
