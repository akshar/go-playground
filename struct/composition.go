package main

import "fmt"

type person struct {
	first string
	last  string
	id    int
}

///////////////////////////////////////////
// (p person) is the receiver		 //
// it is another parameter		 //
// not idiomatic to use "this" or "self" //
///////////////////////////////////////////
func (p person) fullName() string {
	return p.first + p.last
}

func main() {
	p1 := person{"foo", "bar", 1}
	p2 := person{"foo1", "bar2", 2}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}

//https://www.ardanlabs.com/blog/2015/09/composition-with-go.html
