package main

import "fmt"

type Person struct {
	first string
	id    int
}

type foo struct {
	Person
	bar bool
}

func (p Person) greeting() {
	fmt.Println("hello" + p.first)
}

func (p foo) greeting() {
	fmt.Println("hello foo " + p.first)
}

func main() {
	p1 := foo{
		Person: Person{
			first: "abc",
			id:    314,
		},
		bar: true,
	}
	//fields and methods of the inner-type are promoted to the outer type
	p1.greeting()        // hello foo abc
	p1.Person.greeting() // helloabc
}
