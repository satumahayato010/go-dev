package main

import "fmt"

type Set struct {
	items map[int]struct{}
}

func NewSet() *Set {
	return &Set{items: make(map[int]struct{})}
}

func (s *Set) Add(val int) {
	s.items[val] = struct{}{}
}

func (s *Set) Has(val int) bool {
	_, ok := s.items[val]
	return ok
}

func containsDuplicate(nums []int) bool {
	set := NewSet()

	for _, num := range nums {
		if set.Has(num) {
			return true
		}
		set.Add(num)
		fmt.Println(set)
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 4, 1}
	output := containsDuplicate(nums)
	fmt.Println(output)
}
