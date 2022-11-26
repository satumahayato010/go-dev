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
