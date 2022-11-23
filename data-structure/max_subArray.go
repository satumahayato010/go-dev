package data_structure

/*
Maximum Subarray
整数配列numsが与えられたとき値が最大となるサブ配列を見つけ、その和を返す。

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6
*/

func MaxSubArray(nums []int) int {
	maxSum, curSum := nums[0], nums[0]

	for _, num := range nums[1:] {
		if curSum < 0 {
			curSum = num
		} else {
			curSum += num
		}
		if curSum > maxSum {
			maxSum = curSum
		}
	}
	return maxSum
}
