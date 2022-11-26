package skills

/*
Count Odd Numbers in an Interval Range
二つの負でない整数lowとhighが与えられる。lowとhighの間の奇数の数を返せ（この値を含む）

Input: low = 3, high = 7
Output: 3
Explanation: The odd numbers between 3 and 7 are [3,5,7]
*/

func CountOdds(low, high int) int {
	return (high+1)/2 - low/2
}

/*
0から高までの奇数の数は、(高+1) / 2で求めることができ、
0から低までの奇数の数は、低 / 2で求めることができます。
*/
