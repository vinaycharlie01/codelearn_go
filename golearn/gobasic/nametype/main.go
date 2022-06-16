package main

import "fmt"

// //Declare a new type type called duration. Have the underlying type be an int.
//type duration int

//Declare two defined types called mile and kilometer. Have the underlying type be an float64.

type mile float64
type kilometers float64

//Declare a constant called m2km equals 1.609  ( 1mile=1.609km)

const (
	m2km = 1.609
)

func main() {
	// //Coding Exercise #1

	// //Declare a variable of the new type called hour using the var keyword
	// var hours duration
	// //print out the value of the variable hour
	// fmt.Println(hours)
	// //print out the type of the variable hour
	// fmt.Printf("type: %T\n", hours)
	// //assign 3600 to the variable hour using the  = operator
	// hours = 3600
	// //print out the value of hour
	// fmt.Println(hours)

	// //Coding Exercise #2

	// //There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.
	// var hour duration = 3600
	// minute := 60
	// fmt.Println(hour != duration(minute))

	var mileBerlinToParis mile = 655.3
	var kmBerlinToParis kilometers
	kmBerlinToParis = kilometers(mileBerlinToParis * m2km)
	fmt.Printf(kmBerlinToParis)

}
