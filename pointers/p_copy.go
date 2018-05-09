package main

import "fmt"

func foo(x int) {
	x = 0
}

func main() {
	x := 5
	foo(10)
	fmt.Println(x) // still 5
}
