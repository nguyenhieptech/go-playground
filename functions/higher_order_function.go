package main

import "fmt"

func operate(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func HigherOrderFunction() {
	add := func(a, b int) int { return a + b }
	fmt.Println("Sum:", operate(2, 3, add))
}
