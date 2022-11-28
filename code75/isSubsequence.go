package code75

/*
 Is Subsequence
2つの文字列sとtが与えられたとき、sがtの部分配列であればtrueを、そうでなければfalseを返す。
文字列の部分配列とは、元の文字列から、残りの文字の相対的な位置を乱すことなく、
いくつかの文字（なくてもよい）を削除してできた新しい文字列のことである。
(例: "ace" は "abcde" の部分文字列であるが、"aec" は部分文字列ではない)

Input: s = "abc", t = "ahbgdc"
Output: true
*/

func IsSubsequence(s, t string) bool {
	searchIdx := 0
	if s == "" {
		return true
	}

	for _, chr := range t {
		if s[searchIdx] == byte(chr) {
			searchIdx += 1
		}
		if len(s) == searchIdx {
			return true
		}
	}
	return false
}
