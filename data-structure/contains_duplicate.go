package data_structure

/*
Contains Duplicate(重複を含む)
整数の配列 nums が与えられたとき、いずれかの値が配列中に少なくとも2回出現すれば真を、
すべての要素が異なる場合は偽を返す。

Input: nums = [1,2,3,1]
Output: true
*/

func ContainsDuplicate(nums []int) bool {
	m := make(map[int]bool)

	for _, num := range nums {
		if m[num] {
			return true
		}
		m[num] = true
	}
	return false
}
