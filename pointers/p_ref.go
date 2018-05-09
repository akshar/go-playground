package main

import "fmt"

func foo(y int, x *int) {
	*x = y
}

func main() {
	x := 5
	foo(10, &x)
	fmt.Println(x) // 10
}
