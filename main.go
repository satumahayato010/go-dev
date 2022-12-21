package main

import (
	"fmt"
	"go-dev/code75"
)

func main() {
	root := &code75.Node{
		Val: 1,
		Children: []*code75.Node{
			&code75.Node{Val: 3, Children: []*code75.Node{
				&code75.Node{Val: 5},
				&code75.Node{Val: 6},
			}},
			&code75.Node{Val: 2},
			&code75.Node{Val: 4},
		},
	}

	result := code75.Preorder(root)
	fmt.Println(result)
}
