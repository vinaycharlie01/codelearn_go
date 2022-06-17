package main

import "fmt"

/*Coding Exercise #1
Create a function called cube() that takes
a parameter of type float64 and returns the
cube of that parameter (the parameter to the power
 of 3).*/
// func cube(a float64) float64 {
// 	return math.Pow(a, 3)

// }

// /*Coding Exercise #2
// Create a Go program with a function
// called f1() that takes a parameter of type
// uint and returns 2 values:*/
// func f1(s uint) (a, b int) {
// 	//a) the factorial of n
// 	fact := 1
// 	//b) the sum of all integer numbers greater than zero (>0) a
// 	//nd less than or equal to n (<=n)
// 	sum := 0
// 	for i := 1; i <= int(s); i++ {
// 		fact *= i
// 		sum += i
// 	}
// 	return fact, sum
// }

// /*Coding Exercise #3
// Write a function called myFunc()
// that takes exactly one argument which
// is an int number written between double
// quotes (this is in fact a string). If the
// argument is integer 'n', the function should return
// the result of n + nn + nnn*/
// func myFunc(a string) (b string) {
// 	// c := strings.Repeat(a, 2)
// 	// d := strings.Repeat(a, 3)
// 	// e := "+"
// 	b = a + "+" + strings.Repeat(a, 2) + "+" + strings.Repeat(a, 3)
// 	return b

// }

// /*Coding Exercise #4
// Create a function with the identifier sum
//  that takes in a variadic parameter of type
//  int and returns the
// sum of all values of type int passed in.*/
// func varadic(a ...int) int {
// 	sum := 0
// 	for _, v := range a {
// 		sum += v
// 	}
// 	return sum

// }

/*Coding Exercise #5
Change the function from the previous exercise
and use a `naked return`.*/
// func varadic(a ...int) c int {
// 	sum := 0
// 	for _, v := range a {
// 		sum += v
// 	}
// 	return sum

// }

// /*Coding Exercise #6
// Create a function called searchItem() that takes 2
// parameters: a) a string slice and b) a string.
// The function should search for the string
// (the second parameter) in the slice (the first
// parameter) and returns true if it finds the string in the slice and false otherwise.
// The function does a case-sensitive search.*/
// func searchitem(a []string, b string) (c bool) {
// 	for _, v := range a {
// 		if v == b {
// 			return true
// 		}
// 	}
// 	return
// }

// // Coding Exercise #7
// func searchItem(mySlice []string, myStr string) bool {
// 	for _, v := range mySlice {
// 		if strings.EqualFold(v, myStr) {
// 			return true
// 		}
// 	}
// 	return false
// }

// /*Coding Exercise #8
// Consider the following Go program that prints out:
// The Go gopher is the iconic mascot of the Go project.
// Hello, Go playground!
// package main
// import "fmt"
// func print(msg string) {
// 	fmt.Println(msg)
// }
// func main() {
// 	print("The Go gopher is the iconic mascot of the Go project.")
// 	fmt.Println("Hello, Go playground!")
// }
// Modify only the line in the main() body function where the print() function is invoked so that the program will print out Hello, Go playground! and then The Go gopher is the iconic mascot of the Go project.*/
// func print(msg string) {
// 	fmt.Println(msg)
// }

/*Coding Exercise #9
Create a function that takes in an int value
and prints out that value.
Assign the function to a variable, print out the
type of the variable and then call that function
through the
variable name.*/
func foo(a int) {
	fmt.Println(a)
	fmt.Printf("%T", a)

}

func main() {
	///*Coding Exercise #1
	// result := cube(3)
	// fmt.Println(result)

	// a, b := f1(4)
	// fmt.Println("factoril: ", a)
	// fmt.Println("sum: ", b)

	// x := "my new text is this long"
	// y := strings.Repeat("#", len(x))
	// fmt.Println(y)
	// fmt.Printf("%T", y)

	// //Coding Exercise #3
	// fmt.Print("enter the string: ")
	// var a string
	// fmt.Scanln(&a)

	// b := myFunc(a)
	// fmt.Println(b)

	// //Coding Exercise #4
	// a := varadic(10, 20, 30, 40)
	// b := []int{20, 30, 30, 50}
	// d := varadic(b...)
	// fmt.Println(a, d)

	// //Coding Exercise #5
	// a := varadic(10, 20, 30, 40)
	// b := []int{20, 30, 30, 50}
	// d := varadic(b...)
	// fmt.Println(a, d)

	// //Coding Exercise #6
	// animals := []string{"lion", "tiger", "bear"}
	// a := searchitem(animals, "pig")
	// fmt.Println(a)

	// //Coding Exercise #7
	// animals := []string{"lion", "tiger", "bear"}

	// result := searchItem(animals, "BEAR")
	// fmt.Println(result) // => true

	// result = searchItem(animals, "lION")
	// fmt.Println(result) // => true

	// result = searchItem(animals, "pig")
	// fmt.Println(result) // => false

	// //Coding Exercise #8
	// defer print("The Go gopher is the iconic mascot of the Go project.")
	// fmt.Println("Hello, Go playground!")

	// //Coding Exercise #9
	// a := foo
	// fmt.Printf("%T\n", a)
	// a(10)
}
