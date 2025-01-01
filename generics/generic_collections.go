package main

import "fmt"

func printMap[K comparable, V any](m map[K]V) {
	for k, v := range m {
		fmt.Println(k, ":", v)
	}
}

func reverseSlice[T any](s []T) []T {
	n := len(s)
	reversed := make([]T, n)
	for i, v := range s {
		reversed[n-i-1] = v
	}
	return reversed
}

func main() {
	m := map[string]int{"Alice": 30, "Bob": 25}
	printMap(m)

	fmt.Println(reverseSlice([]int{1, 2, 3, 4}))
	fmt.Println(reverseSlice([]string{"A", "B", "C", "D"}))
}
