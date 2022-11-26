package code75

import "fmt"

/*
 Running Sum of 1d Array
配列numsが与えられたとき配列のランニングサムを
runningSum[i] = sum(nums[0]...nums[i])と定義します．
numsの総和を返します。

Input: nums = [1,2,3,4]
Output: [1,3,6,10]
Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].
*/

func RunningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
		fmt.Println(nums)
	}
	return nums
}
