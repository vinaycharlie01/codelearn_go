package main

import "fmt"

func main() {
	// //Coding Exercise #1
	// //1. Declare a map called m1 which value is nil. Print out its type and value.
	// var m1 map[int]string
	// fmt.Println(m1)
	// fmt.Printf("%T\n", m1)
	// //2. Declare a map called m2. It's keys are of type int and values of type string.  Initialize the map using  a map literal with two key:value pairs.
	// m2 := map[int]string{1: "vinay", 2: "kumar"}
	// //3. Add the following key: value to the map: 10: "Abba"
	// m2[10] = "vinay"
	// //4. Retrieve the value of an existing key and the value of a non existing key. Think about the results.

	// fmt.Println(m2[2])
	// fmt.Println(m2[100]) //non

	// /*Coding Exercise #2

	// There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.

	// package main

	// import "fmt"

	// func main() {
	// 	var m1 map[int]bool
	// 	m1[5] = true

	// 	m2 := map[int]int{3: 10, 4: 40}
	// 	m3 := map[int]int{3: 10, 4: 40}

	// 	fmt.Println(m2 == m3)
	// }*/
	// var m1 map[int]bool

	// m2 := map[int]int{3: 10, 4: 40}
	// m3 := map[int]int{3: 10, 4: 40}

	// _, _, _ = m1, m2, m3

	//Coding Exercise #3
	//Consider the following map declaration:
	//m := map[int]bool{"1": true, 2: false, 3: false}
	//1. The above map declaration has an error. Solve the error!

	m := map[int]bool{1: true, 2: false, 3: false}
	//2. Delete a key:value pair from the map.
	delete(m, 1)
	//3. Iterate over the map and print out each key and value.
	for i, v := range m {
		fmt.Println("keys: ", i, "value:", v)

	}

}
