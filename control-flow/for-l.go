package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//while
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	i := 0
	for {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i >= 50 {
			break
		}

	}

}
