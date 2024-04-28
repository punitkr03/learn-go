package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch cases!!")

	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You can open.")
	case 2:
		fmt.Println("Move 2 spot.")
	case 3:
		fmt.Println("Move 3 spot.")
		// fallthrough
		//fallthrough should be explicity added for the expected behaviour
	case 4:
		fmt.Println("Move 4 spot.")
	case 5:
		fmt.Println("Move 5 spot.")
	case 6:
		fmt.Println("Move 6 spot and take one more chance.")
	default:
		fmt.Println("Invalid")
	}
}
