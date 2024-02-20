package main

import "fmt"

func main() {
	var user string = "Punit_Kr"
	fmt.Println("User: ", user)
	fmt.Printf("The type of user is: %T\n", user)

	var isBool bool = true
	fmt.Println("isBool: ", isBool)
	fmt.Printf("The type of isBool is: %T\n", isBool)

	var smallValue uint8 = 255
	fmt.Println("smallValue: ", smallValue)
	fmt.Printf("The type of smallValue is: %T\n", smallValue)

	var smallFloat float32 = 255.255
	fmt.Println("smallValue: ", smallFloat)
	fmt.Printf("The type of smallFloat is: %T\n", smallFloat)

}
