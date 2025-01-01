package main

import "fmt"

func getFirst[T any](items []T) T {
	return items[0]
}

func GenericFunction() {
	fmt.Println(getFirst([]int{1, 2, 3}))          // 1
	fmt.Println(getFirst([]string{"A", "B", "C"})) // A
}
