package main

import "fmt"

type foo int

func main() {

	var x foo
	x = 1
	fmt.Printf("%T %v \n", x, x) // main.foo 1
	var y int
	y = 2
	fmt.Printf("%T %v\n", y, y) // int 2
	fmt.Println(x + y)          //invalid operation: x + y (mismatched types foo and int)
	// This conversion works
	fmt.Println(int(x) + y) // 3

}
