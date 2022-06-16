package main

import "fmt"

func main() {
	// //Coding Exercise #1
	// //1. Using the var keyword declare a string called name and initialize it with your name.
	// var name string = "VINAYKUMAR"
	// //2. Using short declaration syntax declare a string called country and assign the country you are living in to the string variable.
	// country := "INDIA"
	// //3. Print the following string on multiple lines like this:
	// //Your name: `here the value of name variable`
	// //Country: `here the value of country variable`
	// fmt.Println("Your name: ", name)
	// fmt.Println("Country: ", country)
	// //4. Print out the following strings:
	// //a) He says: "Hello"
	// //b) C:\Users\a.txt
	// fmt.Println(`He says: "Hello"`)
	// fmt.Println("He says: \"Hello\"")
	// fmt.Println("C:\\Users\\a.txt")

	// //Coding Exercise #2
	// //1. Declare a rune called r that stores the non-ascii letter ă
	// var a rune = 'ă'
	// //2. Print out the type of r
	// fmt.Printf("%T%v\n", a, string(a))
	// //3. Declare 2 strings that contains the values ma and m
	// var b, c = "ma", "m"
	// //4. Concatenate the strings and the rune in a new string (the new string will hold the value mamă ).
	// concat := b + c + string(a)
	// fmt.Println(concat)
	// /*Coding Exercise #3
	// Consider the following string declaration: s1 := "țară means country in Romanian"*/
	// s1 := "țară means country in Romanian"
	// a := []byte(s1)
	// //1. Iterate over the string and print out byte by byte
	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(a[i])
	// }
	// //2. Iterate  over the string and print out rune by rune and the byte index where the rune starts in the string
	// b := []rune(s1)
	// for _, v := range b {
	// 	fmt.Print(string(rune(v)))
	// }

	// /*Coding Exercise #4
	// There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.
	// package main
	// import "fmt"
	// func main() {
	// 	s1 := "Go is cool!"
	// 	x := s1[0]
	// 	fmt.Println(x)
	// 	s1[0] = 'x'
	// 	// printing the number of runes in the string
	// 	fmt.Println(len(s1))
	// }*/

	// s1 := "Go is cool!"
	// x := s1[0]
	// fmt.Println(x)

	// //s1[0] = "x"

	// // printing the number of runes in the string
	// //fmt.Println(len(s1))
	// fmt.Println(utf8.RuneCountInString(s1))

	// //Coding Exercise #5
	// //Consider the following string declaration:s := "你好 Go!"
	// s := "你好 Go!"
	// //1. Convert the string to a byte slice.
	// d := []byte(s)
	// //2. Print out the byte slice
	// fmt.Println(d)
	// //3. Iterate over the byte slice and print out each index and byte in the byte slice
	// for i, v := range d {
	// 	fmt.Println(i, v)

	// }

	//Coding Exercise #6
	//Consider the following string declaration:s := "你好 Go!"
	s := "你好 Go!"
	//1. Convert the string to a byte slice.
	d := []rune(s)
	//2. Print out the byte slice
	fmt.Println(d)
	fmt.Printf("%#v\n", d)
	//3. Iterate over the byte slice and print out each index and byte in the byte slice
	for i, v := range d {
		fmt.Println(i, string(v))

	}
}
