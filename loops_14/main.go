package main

import "fmt"

func main() {
	fmt.Println("Loops in go")

	days := []string{"Sunday", "Monday", "Tuesday"}
	fmt.Println(days)

	// for d:=0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("Index is %v and value is %v\n", index, day)
	}
}
