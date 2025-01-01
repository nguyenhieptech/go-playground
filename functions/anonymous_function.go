package main

import "fmt"

func AnonymousFunction() {
	double := func(x int) int {
		return x * 2
	}
	fmt.Println("Double of 5:", double(5))
}
