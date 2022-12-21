package code75

/*
Binary Tree Level Order Traversal
二分木のルートが与えられたとき、そのノードの値のレベルオーダーのトラバースを返す。
(すなわち、左から右へ、レベルごとに）

Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	currentLevel := []*TreeNode{root}

	for len(currentLevel) > 0 {
		var nextLevel []*TreeNode
		var values []int

		for _, node := range currentLevel {
			values = append(values, node.Val)
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		result = append(result, values)
		currentLevel = nextLevel
	}
	return result
}
