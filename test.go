package main

import "fmt"

func main() {
	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	m := make(map[string]int)
	fmt.Printf("%T\n", m)

	// ポインタを返す場合は、newを使用する。
	var p *int = new(int)
	fmt.Printf("%T\n", p)

	var st = new(struct{})
	fmt.Printf("%T\n", st)

	// 出力
	// []int
	// map[string]int
	// *int
	// *struct {}
}
