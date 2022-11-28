package code75

/*
Isomorphic Strings
2つの文字列sとtがあるとき、それらが同型であるかどうかを判定せよ。
2つの文字列sとtは、sの文字を置き換えてtを得ることができる場合、同型である。
ある文字の出現箇所はすべて、文字の順序を保ったまま別の文字に置き換えなければならない。
2つの文字が同じ文字に対応することはできないが、1つの文字がそれ自身に対応することは可能である。

Input: s = "egg", t = "add"
Output: true
*/

func IsIsomorphic(s, t string) bool {
	sPat, tPat := map[uint8]int{}, map[uint8]int{}

	for idx := range s {
		if sPat[s[idx]] != tPat[t[idx]] {
			return false
		} else {
			sPat[s[idx]] = idx + 1
			tPat[t[idx]] = idx + 1
		}
	}
	return true
}
