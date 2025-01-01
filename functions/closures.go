package main

import "fmt"

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func Closures() {
	next := counter()
	fmt.Println(next()) // 1
	fmt.Println(next()) // 2
}
