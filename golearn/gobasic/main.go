package main

import "fmt"

func main() {
	//1.Using the var keyword, declare 4 variables called a, b, c, d of type int, float64, bool and string.
	var a int
	var b float64
	var c bool
	var d string
	//2.Using short declaration syntax declare and assign these values to variables x, y and z:
	x := 20
	y := 15.5
	z := "gopher!.."
	//3.Using fmt.Println() print out the values of a, b, c, d, x, y and z.
	fmt.Println(a, b, c, d, x, y, z)
}
