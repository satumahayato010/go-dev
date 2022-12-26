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

/*
logic
基本的なアプローチは、配列を巡回し、現在の部分配列の合計を計算し、それが以前に見つかった最大の合計よりも大きい場合は、最大の合計を更新します。
現在の部分配列の合計を計算する方法は、以下のようになっています。
もし、現在の部分配列の合計が負の場合、それ以前の数字の影響を受けていることがわかるので、現在の数字で部分配列をリセットする。
もし、現在の部分配列の合計が非負の場合、現在の数字を加えることで部分配列を拡張する。
最後に、現在の部分配列の合計が以前に見つかった最大の合計よりも大きい場合は、最大の合計を更新する。
このアルゴリズムは、最悪計算時間が O(n) であるため、配列の長さ n に対して線形時間で実行されます。
*/
