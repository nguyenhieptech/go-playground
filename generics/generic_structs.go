package main

import "fmt"

type Pair[T, U any] struct {
	First  T
	Second U
}

func GenericStruct() {
	p := Pair[int, string]{First: 1, Second: "One"}
	fmt.Println(p)
}
