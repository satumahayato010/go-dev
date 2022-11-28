package main

import (
	"fmt"
	"go-dev/code75"
)

func main() {
	s, t := "egg", "add"
	output := code75.IsIsomorphic(s, t)
	fmt.Println(output)
}
