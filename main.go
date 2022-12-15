package main

import (
	"fmt"
	"go-dev/code75"
)

func main() {
	list1 := &code75.ListNode{Val: 1, Next: &code75.ListNode{Val: 2, Next: &code75.ListNode{Val: 4}}}
	list2 := &code75.ListNode{Val: 1, Next: &code75.ListNode{Val: 3, Next: &code75.ListNode{Val: 4}}}
	output := code75.MergeTwoLists(list1, list2)
	fmt.Println(output)
}
