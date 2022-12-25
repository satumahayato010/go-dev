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

/*
logic
1.配列のすべての要素を2乗して、その結果を同じ配列内に格納します
2.numsと同じ長さの新しい配列outArrが作成
3.2つのポインター、leftとrightを定義しますleftポインターは配列の先頭を指し、right ポインターは配列の末尾を指します。
4.配列の末尾から前に向かって繰り返し処理を行います。そのために、 forループを使用して、iをlen(nums) - 1 から 0 までデクリメントします。
5. left ポインターが指す値と right ポインターが指す値を比較します。
6.nums[left] の方が大きい場合、 outArr[i] に nums[left] を代入し、 left ポインターを1つ進めます。そうでなければ、 outArr[i] に nums[right] を代入し、 right ポインターを1つ戻します。
7.最後に、 outArr を返します。これにより、入力された配列を2乗して昇順に並べ直したものが得られます。
*/
