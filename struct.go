package main

import "fmt"

////////////////////////////////////
// // GO is oop			  //
// 				  //
// 				  //
// 1) Encapsulation		  //
// state('fields')		  //
// behaviour ('methods')	  //
// exported/ un-exported	  //
// 				  //
// 2) Reusability		  //
// inheritance ('embedded types') //
// 				  //
// 3) Polymorphism		  //
// interfaces			  //
// 				  //
// 4) Overriding		  //
// "promotion"			  //
////////////////////////////////////

///////////////////////////////////////////////////////////////
// //Traditional oop					     //
// classes						     //
// -- data strucutre describing a type of object	     //
// -- you than create instance/object from class/blue-Printl //
// -- classes hold both:				     //
// ====State / data / fields				     //
// ====Behaviour / methods				     //
// -- public / private					     //
///////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// //Inheritance					 //
// In Go:						 //
// - you don't create classes, you create a type.	 //
// - you don't instantiate, you create a value of a type //
///////////////////////////////////////////////////////////

type Person struct {
	name string
	id   int
}

type foo struct {
	Person
	bar bool
}

func main() {
	p1 := foo{
		Person: Person{
			name: "abc",
			id:   314,
		},
		bar: true,
	}
	//fields and methods of the inner-type are promoted to the outer type
	fmt.Println(p1.name, p1.Person.name) // abc, abc

}
