package skills

/*
Average Salary Excluding the Minimum and Maximum Salary

salary[i]がi番目の従業員の給与である一意の整数の配列が与えられている．
最小と最大の給与を除いた従業員の平均給与を返せ。実際の答えから10-5以内の答えが認められる。

Input: salary = [4000,3000,1000,2000]
Output: 2500.00000
*/

func Average(salary []int) float64 {
	minNum, maxNum := salary[0], salary[0]
	var sumArr int

	for _, val := range salary {
		sumArr += val

		if val > maxNum {
			maxNum = val
		}
		if val < minNum {
			minNum = val
		}
	}
	return float64(sumArr-minNum-maxNum) / float64(len(salary)-2)
}

/*
logic
このコードは、配列内の最小値と最大値を除いた平均値を計算する関数です。
関数 Average は、整数型の配列 salary を受け取り、float64 型の値を返します。
まず、minNum と maxNum に配列の最初の要素を代入しています。sumArr 変数は、配列内のすべての要素を加算するために使用されます。
次に、for文を使用して配列内のすべての要素を繰り返し処理します。各要素について、sumArr 変数に加算します。
また、要素が以前に保存されている最大値よりも大きい場合は、maxNum 変数を更新します。同様に、要素が以前に保存されている最小値よりも小さい場合は、minNum 変数を更新します。
最後に、sumArr 変数から minNum 変数と maxNum 変数を除いた値を、配列の長さから 2 を引いた数で除算します。これにより、最小値と最大値を除いた配列内の要素の平均値が計算されます。
最終的な結果は、float64 型にキャストして返されます。
*/
