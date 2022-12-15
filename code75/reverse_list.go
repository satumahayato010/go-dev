package code75

/*
Reverse Linked List
単一リンクリストの先頭が与えられたら，リストを反転し，反転後のリストを返す。

Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]
*/

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}
