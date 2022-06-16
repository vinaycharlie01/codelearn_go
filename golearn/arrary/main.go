package main

import "fmt"

func main() {
// 	//Coding Exercise #1
// 	var cities1 [2]string
// 	fmt.Println(cities1)
// 	//1. Using the var keyword, declare an array called cities with
// 	var cities = [2]string{}
// 	//Print out the cities array and notice the value of its elements.
// 	fmt.Println(cities)
// 	//Declare an array called grades of type [3]float64 and initialize only the first 2 elements using an array literal.
// 	var grades = [3]float64{20.2, 2.4}
// 	fmt.Println(grades)
// 	//3. Declare an array called taskDone using the ellipsis operator. The elements are of type bool. Print out taskDone.
// 	taskDone := [...]bool{true}
// 	fmt.Println(taskDone)
// 	//4. Iterate over the array called cities using the classical for loop syntax and the len function and print out element by element.
// 	for i := 0; i < len(cities); i++ {
// 		fmt.Println(cities[i])

// 	}
// 	//5. Iterate over grades using the range keyword and print out element by element.

// 	for _, v := range grades {
// 		fmt.Println(v)

// 	}

// 	//Coding Exercise #2
// 	//Consider the following array declaration: nums := [...]int{30, -1, -6, 90, -6}
// 	//Write a Go program that counts the number of positive even numbers in the array.
// 	nums := [...]int{30, -1, -6, 90, -6}
// 	count := 0
// 	for _, v := range nums {
// 		if v > 0 {
// 			count += 1
// 		}
// 	}
// 	fmt.Println(count)

	//Coding Exercise #3
	//There are some errors in the following Go program. Try to identify the errors, change the code, and run the program without errors.
	/* package main

	import "fmt"

	func main() {
	    myArray := [3]float64{1.2, 5.6}

	    myArray[2] = 6

	    a := 10
	    myArray[0] = a

	    myArray[3] = 10.10

	    fmt.Println(myArray)

	}*/

	myArray := [3]float64{1.2, 5.6}

	myArray[2] = 6

	a := 10.
	myArray[0] = a

	myArray[2] = 10.10

	fmt.Println(myArray)

}
