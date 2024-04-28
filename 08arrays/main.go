package main

import "fmt"

func main() {
	fmt.Println("Arrays are here to stay!")

	var fruits [4]string

	fruits[0] = "Apple"
	fruits[1] = "Papaya"
	fruits[3] = "Banana"

	fmt.Println("Fruit list is: ", fruits)
	fmt.Println("Fruit list length is: ", len(fruits))

	var veggies = [3]string{"Cucumber", "Potato", "Carrot"}
	fmt.Println("Veggie list is: ", veggies)
}
