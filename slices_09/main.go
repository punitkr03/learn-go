package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices!!")

	var fruitList = []string{"Apple", "Peach", "Avacado"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)
	fmt.Println("Fruitlist is : ", fruitList)

	fruitList = append(fruitList, "Mango")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 45
	highScores[1] = 12
	highScores[2] = 76
	highScores[3] = 43

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	//remove an element from an index
	var tempList = []string{"Hello", "Hi", "Mom"}
	fmt.Println(tempList)

	index := 2
	tempList = append(tempList[:index], tempList[index:]...)
	fmt.Println(tempList)
}
