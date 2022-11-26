package algorithm

/*
Search Insert Position
整数のソートされた配列と目標値が与えられたとき、目標が見つかればそのインデックスを返す。
見つからない場合は、順番に挿入された場合のインデックスを返す。

Input: nums = [1,3,5,6], target = 5
Output: 2
*/

func SearchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}
