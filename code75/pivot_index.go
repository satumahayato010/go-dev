package code75

/*
Find Pivot Index
整数numsの配列が与えられたとき、この配列のピボットインデックスを計算しなさい。
ピボットインデックスとは、そのインデックスの左側にあるすべての数値の和が、
そのインデックスの右側にあるすべての数値の和と等しくなるインデックスのことである。
インデックスが配列の左端にある場合、左側には要素がないため、左側の和は0になります。
これは配列の右端にも当てはまります。
左端のピボットインデックスを返します。そのようなインデックスが存在しない場合は、-1 を返します。

Input: nums = [1,2,3]
Output: -1
*/

func PivotIndex(nums []int) int {
	rSum := 0
	for i := 0; i < len(nums); i++ {
		rSum += nums[i]
	}

	lSum := 0
	for i := 0; i < len(nums); i++ {
		if lSum == rSum-nums[i]-lSum {
			return i
		}
		lSum += nums[i]
	}
	return -1
}
