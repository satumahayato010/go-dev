package code75

/*
Validate Binary Search Tree
二分木のルートが与えられたとき、それが有効な二分木探索木（BST）であるかどうかを判定せよ。
有効なBSTは以下のように定義される。
左側のサブツリーには、そのノードのキーより小さいキーを持つノードのみが含まれる。
ノードの右サブツリーには、そのノードのキーより大きいキーを持つノードのみが 含まれる。
左サブツリーと右サブツリーの両方が二分探索木でなければならない。

Input: root = [2,1,3]
Output: true
*/

func IsValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left != nil && root.Left.Val >= root.Val {
		return false
	}

	if root.Right != nil && root.Right.Val <= root.Val {
		return false
	}
	return IsValidBST(root.Left) && IsValidBST(root.Right)
}
