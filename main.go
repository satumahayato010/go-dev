package main

func maxSubArray(nums []int) int {
	maxSum, curSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if curSum < 0 {
			curSum = nums[i]
		} else {
			curSum += nums[i]
		}
		if maxSum < curSum {
			maxSum = curSum
		}
	}
	return maxSum
}

func main() {
}
