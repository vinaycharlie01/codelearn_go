package main

import "fmt"

/*Coding Exercise #1
1. Create your own struct type called person that stores the following data: name, age and a list with favorite colors.
*/

type person struct {
	name      string
	age       int
	favcolors []string
	grades
}

/*Coding Exercise #4

Change the code from Exercise #1 and:
1. Create a new struct type called grades with  2 fields: grade and course*/
type grades struct {
	grade  float64
	cource string
}

func main() {
	// //Coding Exercise #1
	// //2. Declare and initialize two values of type person, one called me and another called you.
	// you := person{
	// 	name:      "john",
	// 	age:       30,
	// 	favcolors: []string{"red", "yello"},
	// }
	// fmt.Println("this is for you:", you)

	// me := person{"vinay", 20, []string{"yello", "bloe"}}
	// //3. Print out the struct values.
	// fmt.Println("this os for me: ", me)
	// me := person{"vinay", 20, []string{"yello", "blow"}}

	// //Coding Exercise #2
	// //1. Change the name or the struct value called me to "Andrei"

	// you := person{"vinay", 20, []string{"yello", "blow"}}

	// //1. Change the name or the struct value called me to "Andrei"
	// you.name = "Andrei"
	// //2. Take in a new variable the favorites colors of struct value called you. Print out the type and the value of the new variable.
	// var colors []string = you.favcolors
	// fmt.Println(colors)
	// //3. Add a new favorite color to the second struct value called you.
	// colors = append(colors, "red", "dark")
	// you.favcolors = colors
	// //4. Print out the struct values.
	// fmt.Println(you)

	// /*Coding Exercise #3
	// Consider the code from Exercise #1.
	// Iterate and print out the favorite colors of the struct value called me.*/

	// you := person{
	// 	name:      "john",
	// 	age:       30,
	// 	favcolors: []string{"red", "yello", "green", "orange"},
	// }

	// for _, v := range you.favcolors {
	// 	fmt.Println(v)
	// }

	a := grades{
		grade:  98.,
		cource: "GOLang",
	}
	//2. Add another field of type grades to person struct type (embedded struct).
	you := person{
		name:      "john",
		age:       30,
		favcolors: []string{"red", "yello", "green", "orange"},
		grades:    a,
	}
	//above
	//3. Change the initialization of the struct values called me and you to contain information about the grades.
	you.grades.cource = "Python"
	you.grades.grade = 98.

	fmt.Println(you)
}
