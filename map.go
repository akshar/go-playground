package main

import "fmt"

// key can be of any type for which equality operator is defined.e.g int, floating point, string, integers, ,structs and array

// slices can't be used as key as equality is not defined on them. Like slices, maps holds references to an underlying Data structure

// attempt to fetch map value with a key that is not present will return the zero value for type of entries in the map
func main() {
	m := make(map[string]int) // or composite literaln:= map[string]int{"foo":1, "bar":2} or map[string]int{}

	m["foo"] = 1
	m["bar"] = 2
	fmt.Println(m)
	fmt.Printf("%T", m)
	value, present := m["foo"]
	fmt.Println()
	fmt.Println(value, present) // 1, true
	delete(m, "bar")
	fmt.Println(m) // map[foo:1]

	map2 := map[string]int{"foo": 1, "bar": 2, "baz": 3}

	if _, exists := map2["baz"]; exists {
		fmt.Println("value exists")
		delete(map2, "foo")
	} else {
		fmt.Println("That value does not exist")
		fmt.Println(map2)
	}

	//Range

	myGreeting := map[int]string{
		0: "Good Morning",
		1: "Bonjour!",
		2: "Bueonos dias!",
	}

	for k, v := range myGreeting {
		fmt.Println(k, " - ", v)
	}
	//https://www.ardanlabs.com/blog/2013/12/macro-view-of-map-internals-in-go.html
	//https://blog.golang.org/go-maps-in-action

}
