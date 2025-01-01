package main

import "fmt"

// When you use multiple defer statements within a function,
// they are executed in a Last-In-First-Out (LIFO) order.
// This means that the most recently deferred function is executed first,
// followed by the second most recently deferred function, and so on.
// https://go.dev/tour/flowcontrol/13

func printSecond() {
	defer fmt.Println("Second")
}

func DeferredFunction() {
	fmt.Println("Hello")

	defer fmt.Println("Third")
	defer printSecond()
	defer fmt.Println("First")

	fmt.Println("Counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done")
}
