package main

//slice type is an abstraction built on top of array
// A slice, once initialized, is always associated with an underlying array that holds its elements
//shares storage with its array and with other slices of same array
import "fmt"

func main() {
	// s := []int{1, 2, 3}
	// fmt.Printf("%T \n", s)
	// fmt.Println(s)

	// or

	// s := make([]int, 5, 6) // type, length, capacity length = len(s)  capacity = cap(s)

	////////////////////////////////////////////////////////////////////
	// make([]int, 50, 100) // new slice of len 50 and cap 100	  //
	// //or								  //
	// new([100]int)[0:50]						  //
	////////////////////////////////////////////////////////////////////

	mySlice := make([]int, 0, 5)
	fmt.Println("-----------------------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("-----------------------------")

	// Doubles the size of the underlying array when len exceeds the capacity
	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), "value:", mySlice[i])
	}

}
