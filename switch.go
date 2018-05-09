package main

import "fmt"

func SwitchOnType(x interface{}) {
	switch x.(type) { //this is an assert; asserting , "x is of this type"
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	name := "foo"
	switch name {
	case "foo":
		fmt.Println("foo")
	case "bar":
		fmt.Println("bar")
	default:
		fmt.Println("foobar")
	}

	//clojure's cond

	// (cond
	// 	(cond1) expr1
	// 	(cond2) expr2
	// 	:else :defaultexpr)
	switch {
	case false:
		fmt.Println("foo")
	case 4%2 == 0: // return from here
		fmt.Println("bar")
	default:
		fmt.Println("Default")
	}

	SwitchOnType(7)
	SwitchOnType("foo")

}
