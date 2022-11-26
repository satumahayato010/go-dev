package data_structure

import (
	"fmt"
	"sort"
)

/*
Merge Sorted Array

非減少順に並べられた2つの整数配列nums1とnums2、
およびnums1とnums2のそれぞれの要素数を表す2つの整数mとnが与えられている。
nums1とnums2を非減少順に並べた1つの配列に統合しなさい．
最終的にソートされた配列は関数から返されず、配列 nums1 の内部に格納されます。
このため、nums1 は m + n の長さを持ち、最初の m 個の要素はマージされる要素を表し、
最後の n 個の要素は 0 に設定され、無視されることになります。

Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
*/

func Merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:n], nums2...)

	sort.Ints(nums1)
}

func Merge2(nums1 []int, m int, nums2 []int, n int) {
	fmt.Println(nums1, nums2)
	for n > 0 {
		if m == 0 || nums1[m-1] <= nums2[n-1] {
			nums1[m+n-1] = nums2[n-1]
			fmt.Println(nums1, nums2)
			n -= 1
		} else {
			nums1[m+n-1] = nums1[m-1]
			fmt.Println(nums1, nums2, 1)
			m -= 1
		}
	}
}
