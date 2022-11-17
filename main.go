package main

import "fmt"

func euclideanAlgorithm(m, n int) int {
	if n == 0 {
		return m
	} else {
		return euclideanAlgorithm(n, m%n)
	}
}

func main() {
	output := euclideanAlgorithm(128, 72)
	fmt.Println(output)
}
