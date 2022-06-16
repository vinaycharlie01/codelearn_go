package main

func main() {
	//Coding Exercise #3
	//There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.

	var a float64 = 7.1

	x, y := true, 3.7

	a, x = 5.5, false
	a = 2.31
	x = false

	_, _, _ = a, x, y
}
