package algorithm

/*
Rotate Array
配列が与えられたとき，配列を右にkステップ回転させる（kは非負）

Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
*/

func Rotate(nums []int, k int) {
	n := len(nums)
	k %= n

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start += 1
		end -= 1
	}
}

/*
logic
Rotate 関数では、以下のような処理が行われています。
1.配列の長さ n を取得し、k を n で割った余りを新しい k として設定することで、k を n 以下の数値に制限します。
2.reverse 関数を使って、配列全体を反転させます。
3.reverse 関数を使って、配列の最初から k-1 番目までを反転させます。
4.reverse 関数を使って、配列の k 番目から最後までを反転させます。
これにより、配列を k回転させることができます。

reverse 関数では、配列の最初から最後までを、交換を繰り返しながら反転させる処理が行われています。 以下のような流れです。
1.start 変数と end 変数を使って、配列を指定された範囲で反転させることができるようにします。
2.start 変数を start + 1、end 変数を end - 1 と更新し、次のループで次の要素を処理するようにします。
3.start 変数が end 変数より大きくなるまで、交換を繰り返します。
これで、配列を指定された範囲で反転させることができます。
*/
