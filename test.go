package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"} //This is a slice
	fmt.Println(colors)
	colors = append(colors, "Purple")
	fmt.Println(colors)
	fmt.Println(colors)

	numbers := make([]int, 4)
	numbers[0] = 14
	numbers[1] = 74
	numbers[2] = 15
	numbers[3] = 45
	fmt.Println(numbers)

	numbers = append(numbers, 65)
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)
}
