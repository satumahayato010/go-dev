package main

import (
	"fmt"
	"go-dev/code75"
)

func main() {
	root := &code75.TreeNode{
		Val: 3,
		Left: &code75.TreeNode{
			Val: 9,
		},
		Right: &code75.TreeNode{
			Val: 20,
			Left: &code75.TreeNode{
				Val: 15,
			},
			Right: &code75.TreeNode{
				Val: 7,
			},
		},
	}

	result := code75.LevelOrder(root)
	fmt.Println(result)
}
