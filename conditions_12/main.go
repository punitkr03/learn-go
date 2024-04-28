package main

import "fmt"

func main() {
	fmt.Println("If else in GO")

	i := 23
	var result string

	if i < 25 {
		result = "Less than 25"
	} else if i > 25 {
		result = "More than 25"
	} else {
		result = "Equal"
	}

	fmt.Println(result)
}
