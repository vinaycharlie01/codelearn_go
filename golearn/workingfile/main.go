package main

import (
	"io/ioutil"
	"log"
)

func main() {
	/*Coding Exercise #1
	Create a new file in the current working directory called info.txt and then close the file. If the file cannot be created (no permissions, wrong path etc) then print out the error and stop the program (do error handling).
	*/
	// a, err := os.Create("info.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// a.Close()

	// //checking if file exists
	// _, err := os.Stat("info.txt")
	// if err != nil {
	// 	if os.IsNotExist(err) {
	// 		log.Fatal("this file doest not exisrt")

	// 	}
	// }

	// // renaming the file
	// err = os.Rename("info.txt", "data.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// //remove the file
	// err := os.Remove("aa.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	/*Coding Exercise #4

	Create a Go Program that reads the entire contents of a file called info.txt into a string. You can use ioutil.ReadAll() or any other function you want.*/

	// file, err := os.Open("info.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// data, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(data))

	// the file value returned by os.Open() is wrapped in a bufio.Scanner just like a buffered reader.
	// a := bufio.NewScanner(file)

	// reading the whole file line by line:
	// for a.Scan() {
	// 	fmt.Println(a.Text())
	// }

	// err = a.Err()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	bs := []byte("hey heollo how are you")
	err := ioutil.WriteFile("info.txt", bs, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
