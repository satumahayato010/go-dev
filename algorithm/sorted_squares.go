package algorithm

/*
Squares of a Sorted Array
降順でソートされた整数の配列 nums が与えられたとき、降順でソートされた各数値の二乗の配列を返す。

Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]
*/

func SortedSquares(nums []int) []int {
	for idx, num := range nums {
		nums[idx] = num * num
	}

	outArr := make([]int, len(nums))
	left, right := 0, len(nums)-1

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[left] > nums[right] {
			outArr[i] = nums[left]
			left += 1
		} else {
			outArr[i] = nums[right]
			right -= 1
		}
	}
	return outArr
}
