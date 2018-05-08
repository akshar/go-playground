package main

import "fmt"

const (
	Pi       = 3.14
	Language = "GO"
)

const (
	catA = iota //0
	catB        //1
	catC        //2
)

const (
	_  = iota             // ignore 0
	KB = 1 << (iota * 10) // shift 10 times
	MB = 1 << (iota * 10) // shfit 20 times..
	GB = 1 << (iota * 10)
	TB = 1 << (iota * 10)
)

func main() {
	const q = "foobar"
	fmt.Println(q)
	fmt.Printf("%T", q)

	fmt.Printf("%d \t", KB) // %d decimal %b binary
	fmt.Printf("%d \n", MB)
	fmt.Printf("%d \t", GB)
	fmt.Printf("%d \n", TB)
	fmt.Printf("%b \t", KB)
	fmt.Printf("%b \n", MB)
	fmt.Printf("%b \t", GB)
	fmt.Printf("%b \t", TB)
}

//https://blog.golang.org/constants READ THIS!
