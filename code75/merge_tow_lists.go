package code75

/*
Merge Two Sorted Lists
2つのソートされたリンクリスト list1 と list2 の先頭が与えられる．
この2つのリストを1つのソートされたリストに統合せよ。
このリストは最初の2つのリストのノードをつなぎ合わせて作らなければならない。
マージされたリンクリストの先頭を返せ。

Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res

	for list1 != nil || list2 != nil {
		if list1 == nil {
			cur.Next = list2
			list2 = nil
			continue
		}
		if list2 == nil {
			cur.Next = list1
			list1 = nil
			continue
		}
		if list1.Val > list2.Val {
			cur.Next = list2
			cur, list2 = cur.Next, list2.Next
		} else {
			cur.Next = list1
			cur, list1 = cur.Next, list1.Next
		}
	}
	return res.Next
}
