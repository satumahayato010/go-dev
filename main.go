package main

import (
	"fmt"
	"go-dev/algorithm"
)

const personName = "tanaka"

func getPersonName() {
	fmt.Println(personName)

	if true {
		fmt.Println(personName)
	}
}
func main() {
	nums := []int{-4, -1, 0, 3, 10}
	output := algorithm.SortedSquares(nums)
	fmt.Println(output)

	fmt.Println(personName)
	getPersonName()

}
