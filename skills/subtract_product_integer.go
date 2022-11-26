package skills

import "fmt"

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
		fmt.Println(mod)
		sum += mod
		product *= mod

		n /= 10
		fmt.Println(n)
	}
	return product - sum
}
