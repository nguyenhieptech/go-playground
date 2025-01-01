package main

import "fmt"

func doubleNum(value int, callback func(int)) {
	callback(value * 2)
}

func CallbackFunction() {
	doubleNum(5, func(result int) {
		fmt.Println("Result:", result)
	})
}
