package main

import "fmt"

func main() {
	defer fmt.Println("Working ?")
	defer fmt.Println("Two ?")
	defer fmt.Println("One ?")
	fmt.Println("Defer keyword")
}

//Defer follows LIFO property