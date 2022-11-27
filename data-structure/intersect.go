package data_structure

/*
Intersection of Two Arrays
2つの整数配列 nums1 と nums2 が与えられたとき、それらの交点を配列として返す。
結果の各要素は両方の配列に現れる回数だけ出現しなければならず、結果をどのような順序で返してもよい。

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]
*/

func Intersect(nums1, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m := map[int]int{}
	for _, v := range nums1 {
		m[v] += 1
	}

	var resultArr []int
	for _, v := range nums2 {
		if m[v] != 0 {
			resultArr = append(resultArr, v)
			m[v] -= 1
		}
	}
	return resultArr
}
