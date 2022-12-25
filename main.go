package main

import (
	"fmt"
	"go-dev/code75"
)

func main() {
	root := &code75.TreeNode{
		Val: 2,
		Left: &code75.TreeNode{
			Val: 1,
		},
		Right: &code75.TreeNode{
			Val: 3,
		},
	}
	result := code75.IsValidBST(root)
	fmt.Println(result)
}
