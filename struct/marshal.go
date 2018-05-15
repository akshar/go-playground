package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	id    int
}

func main() {
	p1 := person{"foo", "bar", 1}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
}
