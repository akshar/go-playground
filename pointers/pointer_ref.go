package main

import "fmt"

func main() {
	a := 314
	fmt.Println(a)  // 314
	fmt.Println(&a) //0xc4200160d0

	var b *int = &a // b is of type "Int pointer"
	fmt.Println(b)  //0xc4200160d0
	fmt.Println(*b) // 314  deref

	*b = 42
	fmt.Println(a) //42

}
