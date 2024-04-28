package main

import "fmt"

func main() {
	fmt.Println("Pointers are here to stay!")

	// var ptr *int
	// fmt.Println("Value of ptr is: ", ptr)

	Number := 23

	var ptr *int = &Number
	fmt.Println("Value of pointer is: ", ptr)
	fmt.Println("Value of Number is: ", *ptr)
}
