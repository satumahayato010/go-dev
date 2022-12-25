package data_structure

/*
Contains Duplicate(重複を含む)
整数の配列 nums が与えられたとき、いずれかの値が配列中に少なくとも2回出現すれば真を、
すべての要素が異なる場合は偽を返す。

Input: nums = [1,2,3,1]
Output: true
*/

func ContainsDuplicate(nums []int) bool {
	m := make(map[int]bool)

	for _, num := range nums {
		if m[num] {
			return true
		}
		m[num] = true
	}
	return false
}

/*
logic
このコードは、配列 nums に重複した数字があるかどうかを調べる関数です。
1.空のマップ m を作成します。このマップは数字をキーとし、値として true もしくは false を保持します。
2.for ループを使用して、配列 nums の各要素 num を反復処理する.マップ m に num がキーとして登録されているかどうかを調べます。
3.もし、m[num] の値が true であれば、関数は true を返して処理を終了します。
4.もし、m[num] の値が false であれば、マップに num をキーとして、値を true として登録します。
5.配列 nums の全ての数字を処理した後、関数は false を返して処理を終了します。

このアルゴリズムでは、マップ m を使用して、配列 nums のすべての要素を確認します。そして、要素がマップにすでに含まれている場合は、その要素は重複していると判断して true を返します。
もし、すべての要素がマップに含まれていない場合は、配列 nums に重複する要素はないと判断して、false を返します。
*/
