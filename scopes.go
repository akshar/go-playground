package main

import (
	"fmt"

	"./scopedemo"
)

// package level scope
var x int = 42

func main() {
	fmt.Println(scopedemo.Fullname())
}
