package skills

/*
Subtract the Product and Sum of Digits of an Integer
整数nが与えられたとき、その桁の積と桁の和の差を返す。

Input: n = 234
Output: 15
Explanation:
Product of digits = 2 * 3 * 4 = 24
Sum of digits = 2 + 3 + 4 = 9
Result = 24 - 9 = 15
*/

func SubtractProductAndSum(n int) int {
	sum, product := 0, 1

	for n > 0 {
		mod := n % 10
		sum += mod
		product *= mod

		n /= 10
	}
	return product - sum
}

/*
logic
このプログラムは、整数 n の各桁の数字の総和と積を求めて、その差を返す関数です。
まず、変数 sum と product をそれぞれ 0 と 1 で初期化します。次に、整数 n が 0 より大きい間、以下の処理を繰り返します。
1.整数 n を 10 で割ったあまり (mod) を取得します。これにより、整数 n の最下位桁の数字が取り出されます。
2.取り出した数字を変数 sum に加算します。
3.取り出した数字を変数 product に乗算します。
4.整数 n を 10 で割り、その結果を整数型で再代入します。これにより、整数 n の最下位桁が削除され、次に上の桁の数字が取り出されるようになります。
最後に、プログラムは、変数 product の値から変数 sum の値を引いた差を返します。
*/
