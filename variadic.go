package main

import "fmt"

func main() {
	fmt.Println(average(1, 2, 3, 4, 5, 6))
}

func average(nums ...float64) float64 {
	var total float64
	for _, v := range nums {
		total += v
	}

	return total / float64(len(nums))
}
