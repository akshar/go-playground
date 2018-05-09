package main

import "fmt"

// rune: a character, alias for int32, an integer value identifying a unicode code point

// every rune is basically a number
// 4byte : utf-8

func main() {
	fmt.Println([]byte("Hello")) //[72 101 108 108 111] numeric code point
	for i := 50; i <= 150; i++ {
		// fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}

	foo := 'a'
	fmt.Println(foo)            // 97
	fmt.Printf("%T \n", foo)    // int32
	fmt.Printf("%T", rune(foo)) // int32

}
