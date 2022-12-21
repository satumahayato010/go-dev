package code75

import "fmt"

/*
N-ary Tree Preorder Traversal
n-ary木のルートが与えられたとき、そのノードの値の前置トラバーサルを返す。
Nary-Treeの入力の直列化は、そのレベルオーダーのトラバーサルで表現される。
子ノードの各グループはヌル値で区切られる (例参照)

Input: root = [1,null,3,2,4,null,5,6]
Output: [1,3,5,6,2,4]
*/

type Node struct {
	Val      int
	Children []*Node
}

func Preorder(root *Node) []int {
	if root == nil {
		return nil
	}

	result := []int{root.Val}
	for _, child := range root.Children {
		result = append(result, Preorder(child)...)
		fmt.Println(result)
	}

	return result
}
