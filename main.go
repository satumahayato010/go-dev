package main

import (
	"fmt"
	"go-dev/code75"
)

func main() {
	head := &code75.ListNode{Val: 1,
		Next: &code75.ListNode{Val: 2,
			Next: &code75.ListNode{Val: 3,
				Next: &code75.ListNode{Val: 4,
					Next: &code75.ListNode{Val: 5}}}}}
	output := code75.ReverseList(head)
	fmt.Println(output)
}
