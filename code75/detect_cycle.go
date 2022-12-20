package code75

/*
Linked List Cycle II
リンクリストの先頭が与えられたとき、サイクルが開始するノードを返す。
サイクルが存在しない場合は、nullを返す。
リンクリストにサイクルが存在するのは、
リスト内に次のポインタを追い続ければ再び到達できるノードが存在する場合である。
内部的には、posはtailの次のポインタが接続されている
ノードのインデックスを示すために使われる(0-indexed)。サイクルが存在しない場合は-1である。
posはパラメータとして渡されないことに注意。
リンクリストを修正してはならない。

Input: head = [3,2,0,-4], pos = 1
Output: tail connects to node index 1
*/

func detectCycle(head *ListNode) *ListNode {
	// Use the two pointer technique to detect if there is a cycle
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			// There is a cycle, so find the start of the cycle
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}

	// There is no cycle
	return nil
}
