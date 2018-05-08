package main

import "fmt"

// declaration : can only be used inside fn
//var: every unassigned var has zero value int =0 , bool = false, string = ""
func main() {
	//shorthand declare assign
	a := 10
	b := "golang"
	c := 3.14
	d := true

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

	//or

	var x string
	x = "foo"
	fmt.Println(x)

	// or

	var f, z int = 1, 2
	fmt.Println(f, z)

}
