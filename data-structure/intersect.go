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

/*
logic
まず、関数の初めで、nums1とnums2のスライスの長さを比較し、nums1が長い場合は、nums1とnums2を入れ替えます。
これは、後に処理をするときに、長さが短い方を基準にして処理をすることで、処理時間を短縮するためです。
次に、nums1を使用して、map型の変数mを作成します。このmapは、整数をキーとし、その整数がnums1にいくつあるかを値として保持します。

次に、nums2を見ていきます。nums2の各要素vが、mに登録されているとき、つまりnums1にも存在するとき、
その要素を結果の配列resultArrに追加します。また、mからその要素を1つ減らします。
最後に、resultArrを返します。
*/
