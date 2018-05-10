package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int {
	even := []int{}
	for _, n := range numbers {
		if callback(n) {
			even = append(even, n)
		}
	}
	return even
}

func main() {
	xs := filter([]int{1, 2, 3, 4, 5}, func(n int) bool {
		return n%2 == 0
	})

	fmt.Println(xs) // [2,4]
}
