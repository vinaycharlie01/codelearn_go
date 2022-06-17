package main

/*//Coding Exercise #1
// package main
// import (
//     "fmt"
// )
// func sayHello(n string) {
//     fmt.Printf("Hello, %s!\n", n)
// }
// func main() {
//     go sayHello("Mr. Wick")
// }
// Your task is to synchronize main and the goroutine using WaitGroups. The program should print the string received as argument by sayHello().
func sayHello(n string, wg *sync.WaitGroup) {
	//defer wg.Done()
	fmt.Printf("Hello, %s!\n", n)
	wg.Done()
}
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go sayHello("Mr. Wick", &wg)
	wg.Wait()
}*/

/*// Coding Exercise #2
//1. Create a function called sum() that
// calculates and then prints out  the sum
//of 2 float numbers
//it receives as arguments.
func sum(a, b float64, wg *sync.WaitGroup) {
	defer wg.Done()
	add := a + b
	fmt.Printf("%.2f", add)

}
func main() {
	//3. Synchronize the goroutines and the main function using WaitGroups
	var wg sync.WaitGroup
	wg.Add(1)
	//2. From main launch 3 goroutines that execute the function you have just created (sum)
	go sum(10.19199, 20.19929, &wg)
	go sum(10, 20, &wg)
	go sum(10., 20., &wg)
	wg.Wait()

}*/

/*//Coding Exercise #3
func main() {
	//2. Launch the function as a goroutine and synchronize it with main using WaitGroups
	var wg sync.WaitGroup
	wg.Add(1)
	//1. Create an anonymous function that calculates and prints out the square root of a float value it receives as argument.
	go func(a float64, wg *sync.WaitGroup) {
		defer wg.Done()
		x := math.Sqrt(a)
		fmt.Println(x)

	}(16.10, &wg)
	wg.Wait()

}*/

/*//Coding Exercise #4
//Change the code from Exercise #3 and launch 50 goroutines that calculate concurrently the square root of all the numbers between 100 and 149 (both included).
func main() {

	var wg sync.WaitGroup
	wg.Add(50)
	for i := 100.; i < 150.; i++ {
		go func(f float64, wg *sync.WaitGroup) {
			defer wg.Done()
			x := math.Sqrt(f)
			fmt.Println(x)
		}(i, &wg)
	}
	wg.Wait()

}*/

/*Coding Exercise #5

You work at a Banking Application and have created 2 functions: one that deposits a value into an account and another that withdraws a value from the account.

You want to simulate many deposits and withdraws that take place simultaneously  and start some goroutines.

During testing you notice that a date race occurred.

Your task is to change the code in order to protect the account's balance from simultaneously writing using a mutex.

This is the initial program that has errors:

package main

import (
    "fmt"
    "sync"
)

func deposit(b *int, n int, wg *sync.WaitGroup) {
    *b += n
    wg.Done()
}

func withdraw(b *int, n int, wg *sync.WaitGroup) {
    *b -= n
    wg.Done()
}

func main() {
    var wg sync.WaitGroup

    wg.Add(200)

    balance := 100

    for i := 0; i < 100; i++ {
        go deposit(&balance, i, &wg)
        go withdraw(&balance, i, &wg)
    }
    wg.Wait()

    fmt.Println("Final balance value:", balance)
}*/


/*//Coding Exercise #5
func deposit(b *int, n int, wg *sync.WaitGroup, wm *sync.Mutex) {
	wm.Lock()
	*b += n
	wm.Unlock()
	wg.Done()
}

func withdraw(b *int, n int, wg *sync.WaitGroup, wm *sync.Mutex) {
	wm.Lock()
	*b -= n
	wm.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var wm sync.Mutex

	wg.Add(200)

	balance := 100

	for i := 0; i < 100; i++ {
		go deposit(&balance, i, &wg, &wm)
		go withdraw(&balance, i, &wg, &wm)
	}
	wg.Wait()

	fmt.Println("Final balance value:", balance)
}*/
