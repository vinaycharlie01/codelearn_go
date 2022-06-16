package main

import "fmt"

func main() {
	// //Coding Exercise #1
	// //Using a composite literal declare and initialize a slice of type string with 3 elements.
	// slice := []string{"1", "2", "3"}
	// //Iterate over the slice and print each element in the slice and its index.
	// for _, v := range slice {
	// 	fmt.Println(v)
	// }
	// //Coding Exercise #2
	// //There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.
	// /* mySlice := []float64{1.2, 5.6}

	//    mySlice[2] = 6

	//    a := 10
	//    mySlice[0] = a

	//    mySlice[3] = 10.10

	//    mySlice = append(mySlice, a)

	//    fmt.Println(mySlice)*/

	// mySlice := []float64{1.2, 5.6}

	// //mySlice[2] = 6.

	// a := 10.
	// mySlice[0] = a

	// //mySlice[3] = 10.10

	// mySlice = append(mySlice, a)

	// fmt.Println(mySlice)

	// //Coding Exercise #3
	// // 1. Declare a slice called nums with 3 float64 numbers.
	// nums := []float64{1.1, 2.2, 3.3}

	// // 2. Append the value 10.1 to the slice
	// nums = append(nums, 10.1)

	// // 3. In one statement append to the slice the values: 4.1,  5.5 and 6.6
	// nums = append(nums, 4.1, 5.5, 6.6)

	// // 4. Print out the slice
	// fmt.Println(nums)

	// // 5. Declare a slice called n with 2 float64 values
	// n := []float64{10.10, 20.20}

	// // 6. Append n to nums
	// nums = append(nums, n...)

	// // 7. Print out the slice
	// fmt.Println(nums)

	/*Coding Exercise #4
	Create a Go program that reads some numbers from the command line and then calculates the sum and the product of all the numbers given at command line.
	The user should give between 2 and 10 numbers.*/
	// a := []int{3, 2, 5}
	// sum := 0
	// product := 1
	// for _, v := range a {
	// 	sum += v
	// 	product *= v

	// }
	// fmt.Println("Sum: ", sum, ",Product:", product)
	// if len(os.Args) < 2 { //if not run with arguments
	// 	log.Fatal("Please run the program with arguments (2-10 numbers)!")

	// }

	// //taking the arguments in a new slice
	// args := os.Args[1:]

	// // declaring and initializing sum and product of type float64
	// sum, product := 0., 1.

	// if len(args) < 2 || len(args) > 10 {
	// 	fmt.Println("Please enter between 2 and 10 numbers!")
	// } else {

	// 	for _, v := range args {
	// 		num, err := strconv.ParseFloat(v, 64)
	// 		if err != nil {
	// 			continue //if it can't convert string to float64, continue with the next number
	// 		}
	// 		sum += num
	// 		product *= num
	// 	}
	// }

	// fmt.Printf("Sum: %v, Product: %v\n", sum, product)

	// /*Coding Exercise #5
	//   Consider the following slice declaration: nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	// */
	// nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	// //Using a slice expression and a for loop iterate over the slice ignoring the first and the last two elements.
	// slice := nums[len(nums)-len(nums)+1 : len(nums)-2]
	// fmt.Println(slice)

	// sum := 0
	// for _, v := range slice {
	// 	sum += v

	// }
	// //Print those elements and their sum.
	// fmt.Println("sum:", sum)

	// /*Coding Exercise #6
	// Consider the following slice declaration: friends := []string{"Marry", "John", "Paul", "Diana"}
	// Using copy() function create a copy of the slice. Prove that the slices are not connected by modifying one slice and notice that the other slice is not modified.*/
	// nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	// numsCopy := make([]int, len(nums[2:len(nums)-2]))
	// copy(numsCopy, nums)
	// numsCopy = append(numsCopy, 20)
	// fmt.Println("copy:", numsCopy)
	// fmt.Println("not copy", nums)

	// /*Coding Exercise #7
	// Consider the following slice declaration: friends := []string{"Marry", "John", "Paul", "Diana"}*/
	// friends := []string{"Marry", "John", "Paul", "Diana"}
	// yourfriends := []string{}
	// yourfriends = append(yourfriends, friends...)
	// yourfriends[0] = "joan"
	// fmt.Println(friends)
	// fmt.Println(yourfriends)

	/*Coding Exercise #8
	Consider the following slice declaration:
	 years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	Using a slice expression and append() function create a new slice called newYears that contains the first 3 and the last 3 elements of the slice.  newYears should be []int{2000, 2001, 2002, 2008, 2009, 2010}*/
	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	newYears := []int{}
	newYears = append(years[:3], years[len(years)-3:]...)
	fmt.Println(newYears)
}
