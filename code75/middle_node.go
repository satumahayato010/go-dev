package code75

/*
Middle of the Linked List
単一リンクリストの先頭が与えられたとき、そのリンクリストの中央のノードを返す。
中間ノードが2つある場合は、2番目の中間ノードを返す。

Input: head = [1,2,3,4,5]
Output: [3,4,5]
*/

func MiddleNode(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
