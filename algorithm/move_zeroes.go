package algorithm

/*
Move Zeroes
整数配列 nums が与えられたとき、0 以外の要素の相対的な順序を維持したまま、
すべての 0 をその末尾に移動させる。
この処理は、配列のコピーを作成することなく、インプレースで行わなければならないことに注意。

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
*/

func MoveZeroes(nums []int) {
	nonZeroIdx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[nonZeroIdx] = nums[nonZeroIdx], nums[i]
			nonZeroIdx += 1
		}
	}
}

/*
logic
変数 nonZeroIdx を宣言して、その値を 0 に初期化します。
この変数は、現在までに処理された非ゼロ要素の最後の位置を示します。
各要素について、以下の処理を行います。

1.要素がゼロであるかどうかを調べます。もしゼロであれば、何もしません。
もしゼロでなければ、次の処理を行います。

2.変数 i の位置の要素と、変数 nonZeroIdx の位置の要素を交換します。
これにより、非ゼロ要素が前に移動し、ゼロ要素が後ろに移動することになります。

3.変数 nonZeroIdx をインクリメント（加算）します。

*/
