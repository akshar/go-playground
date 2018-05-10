package main

import "fmt"

// name ...

func makeGreeter() func() string {
	return func() string {
		return "FooBar"
	}
}

func main() {
	greeting := func() {
		fmt.Println("foo")
	}

	greeting()

	greeter := makeGreeter()
	fmt.Println(greeter()) //FooBar
}
