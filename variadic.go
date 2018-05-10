package main

import "fmt"

func main() {
	fmt.Println(average(1, 2, 3, 4, 5, 6))

	data := []float64{42,314,11,7}
	fmt.Println(average(data...)) // 93.5
}

func average(nums ...float64) float64 {
	fmt.Println(nums) // [1 2 3 4 5 6] type f64
	var total float64
	for _, v := range nums {
		total += v
	}

	return total / float64(len(nums))
}
