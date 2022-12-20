package code75

/*
Longest Palindrome
小文字または大文字からなる文字列 s が与えられたとき、それらの文字で作ることのできる最長の回文
の長さを返す。
大文字と小文字は区別される。例えば、"Aa "はここでは回文とみなされない。

Input: s = "abccccdd"
Output: 7
*/

func longestPalindrome(s string) int {
	count := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if _, ok := count[ch]; !ok {
			count[ch] = 1
		} else {
			count[ch] += 1
		}
	}

	ans := 0
	for _, val := range count {
		ans += val / 2 * 2
		if ans%2 == 0 && val%2 == 1 {
			ans += 1
		}
	}

	return ans
}
