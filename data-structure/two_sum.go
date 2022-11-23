package data_structure

/*
Two Sum
整数numsの配列と整数targetが与えられたとき、
2つの数値の足し算がtargetになるようなインデックスを返す。
各入力は正確に1つの解を持つと仮定してよく、同じ要素を2度使ってはならない。

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
*/

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for idx, num := range nums {
		if pos, ok := m[target-num]; ok {
			return []int{pos, idx}
		}
		m[num] = idx
	}
	return []int{}
}
