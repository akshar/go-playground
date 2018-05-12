package main

import "fmt"

func main() {
	records := make([][]string, 0)
	// student 1

	student1 := make([]string, 4)
	student1[0] = "Foo"
	student1[1] = "Bar"
	student1[2] = "100"
	student1[3] = "74"

	records = append(records, student1)

	student2 := make([]string, 4)
	student2[0] = "Foo2"
	student2[1] = "Bar2"
	student2[2] = "101"
	student2[3] = "741"

	records = append(records, student2)
	fmt.Println(records)

	//create
	var student []string // or student := []string{} or make([]string,10,20)
	//var student3 [][]string
	// student[0] = "foo"  // out of range, should append
	student = append(student, "foo")
	fmt.Println(student) // ["foo"]

}
