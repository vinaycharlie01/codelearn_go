package main

/*// Coding Exercise #1
// Consider the following type and interface declaration.
// type vehicle interface {
//     License() string
//     Name() string
// }
// type car struct {
//     licenceNo string
//     brand     string
// }

type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenceNo string
	brand     string
}

//1. Create a Go program where car type
// implements the vehicle interface.
func (c car) License() string {
	return c.licenceNo
}

func (c car) Name() string {
	return c.brand

}

// func foo(v vehicle) {
// 	fmt.Println(v.License())
// 	fmt.Println(v.Name())
// }
func main() {
	// a := car{
	// 	brand:     "Tesla",
	// 	licenceNo: "INDIKR007",
	// }

	// foo(a)

	//2. Create a variable of type vehicle that
	// holds a car struct value.
	var a vehicle = car{
		licenceNo: "INDIKR0006",
		brand:     "TATA",
	}
	//3. Call the methods (Licence and Name)
	// on the interface value declared at step 2
	fmt.Println(a.License())
	fmt.Println(a.Name())
	//4. Run the program without errors.
}*/

/*//Coding Exercise #2
func main() {
	//1. Declare a variable called empty of type empty interface.
	// Print out its type.
	var a interface{}
	//2. Assign an int value to the variable called empty. Print out its type.
	a = 10
	fmt.Println(a)
	//3. Assign a float64 value to empty. Print out its type.
	a = 10.2
	fmt.Println(a)
	//4. Assign an int slice value to empty. Print out its type.
	a = []int{10, 20, 30, 0}
	fmt.Println(a)
	//5. Add a new int value to the slice (empty variable).
	a = append(a.([]int), 10)
	//6. Print out the slice (empty variable).
	fmt.Println(a)
}*/

/*//Coding Exercise #3
// package main

// import "fmt"

// type cube struct {
//     edge float64
// }
// func volume(c cube) float64 {
//     return c.edge * c.edge * c.edge
// }

// func main() {
//     var x interface{}
//     x = cube{edge: 5}
//     v := volume(x)
//     fmt.Printf("Cube Volume: %v\n", v)
// }
type cube struct {
	edge float64
}

func volume(c cube) float64 {
	return c.edge * c.edge * c.edge
}
func main() {
	var x interface{}
	x = cube{edge: 5}
	v := volume(x.(cube))
	fmt.Printf("Cube Volume: %v\n", v)
}*/
