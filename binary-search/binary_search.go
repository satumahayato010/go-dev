package binary_search

/*
Binary Search

昇順でソートされた整数の配列numsと整数のtargetが与えられたとき，
numsからtargetを検索する関数を作成せよ．targetが存在する場合、そのインデックスを返す。
そうでなければ -1 を返す。
あなたは，実行時計算量がO(log n)であるアルゴリズムを書かなければならない．

Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4
*/

func LinearSearch(nums []int, target int) int {
	for idx, _ := range nums {
		if nums[idx] == target {
			return idx
		}
	}
	return -1
}

func BinarySearch(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func BinarySearchRecursive(nums []int, target, low, high int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if nums[mid] == target {
		return mid
	} else if nums[mid] < target {
		return BinarySearchRecursive(nums, target, mid+1, high)
	} else {
		return BinarySearchRecursive(nums, target, low, mid-1)
	}
}
