package main

import (
	"fmt"
	"go-dev/code75"
)

func main() {
	s, t := "abc", "ahbgdc"
	output := code75.IsSubsequence(s, t)
	fmt.Println(output)
}
