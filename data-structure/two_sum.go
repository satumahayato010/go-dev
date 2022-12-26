package data_structure

/*
Two Sum
整数numsの配列と整数targetが与えられたとき、
2つの数値の足し算がtargetになるようなインデックスを返す。
各入力は正確に1つの解を持つと仮定してよく、同じ要素を2度使ってはならない。

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
*/

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for idx, num := range nums {
		if pos, ok := m[target-num]; ok {
			return []int{pos, idx}
		}
		m[num] = idx
	}
	return []int{}
}

/*
logic
この関数は、配列 nums の中から 2 つの数を選んでその和が target と等しくなるようなインデックスを返すものです。
アルゴリズムの流れは次のようになっています。
1.空のマップ m を作成する。このマップは数値をキーとし、その数値が nums 配列内で出現する最初のインデックスを値として保存します。
2.nums 配列を range を使って繰り返し処理します。各要素 num とそのインデックス idx を取得します。
3.m マップ内に、target-num の値が存在するかどうかを確認します。もし存在する場合、そのキーに対応する値が
num のインデックス、そして idx が num の対応する値であることから、2 つの数値の和が target と等しいことが分かります。
そのため、このときには []int{pos, idx} を返します。
4.num の値をキー、idx を値として、m マップに保存します。
5.上記の処理を繰り返します。
最終的に、配列 nums の中から 2 つの数を選んでその和が target と等しくなるようなインデックスを見つけることができなかった場合、空の配列を返すようになっています。
*/
