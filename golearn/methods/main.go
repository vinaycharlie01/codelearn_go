package main

import "fmt"

/*//Coding Exercise #1
//1. Create a new type called money.
// Its underlying type is float64.
//2. Create a method called print that
//type memory float64
//prints out the money
//value with exact 2 decimal points.
func (a memory) print() {
	fmt.Printf("%T\n", a)
	fmt.Println(a)
}
func main() {
	a.print()
}*/

/*// Coding Exercise #2

// Consider the money type declared at exercise #1.
//  Create a new method for the money type called printStr that returns the money
//  value as a string with 2 decimal points.
type memory float64
func (a memory) printstr() string {
	return fmt.Sprintf("%.2f", a)
}
func main() {
	var b memory = 10.29349393
	d := b.printstr()
	fmt.Println(d)
}*/

/*//Coding Exercise #3
// 1. Create a new struct type called
// book with 2 fields:
// price and title of type float64 and string.
type book struct {
	price float64
	title string
}
//2. Create a method for the newly defined type called vat
// that returns the vat value if vat is 9%.
func (b book) vat() float64 {
	return b.price * 0.09
}

//3. Cre/ate a method called discount that takes a book value as receiver
//and decreases the price of the book by 10%.
func (b *book) discount() {
	(*b).price = b.price * 0.9
}
func main() {
	a := book{
		price: 40.43,
		title: "golang",
	}
	d := a.vat()
	fmt.Println(d)
	a.discount()
	fmt.Println(a)
}*/

/*Coding Exercise #4

A junior Go Programmer has just developed the following Go Program.   You want the change the book's price by calling changePrice method. However, you notice that after calling the method the price is not changed.

You task is to change the code in order for the changePrice method to work as expected.

package main

import "fmt"

type book struct {
    title string
    price float64
}

// method for book type
func (b book) changePrice(p float64) {
    b.price = p
}
func main() {
    // book value
    bestBook := book{title: "The Trial by Kafka", price: 9.9}

    // changing the price
    bestBook.changePrice(11.99)

    fmt.Printf("Book's price:%.2f\n", bestBook.price) // no change
}
type book struct {
	title string
	price float64
}
// method for book type
func (b *book) changePrice(p float64) {
	(*b).price = p
}
func main() {
	// book value
	bestBook := book{title: "The Trial by Kafka", price: 9.9}

	// changing the price
	bestBook.changePrice(11.99)

	fmt.Printf("Book's price:%.2f\n", bestBook.price) // no change
}*/
