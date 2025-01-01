package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	Position string
}

func StructEmbedding() {
	emp := Employee{
		Person:   Person{Name: "John", Age: 30},
		Position: "Software Engineer",
	}
	fmt.Println(emp.Name, "is a", emp.Position)
}
