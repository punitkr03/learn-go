// pkg.go.dev
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the playground!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")

	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred while reading input. Please try again", err)
		return
	}
	fmt.Println("Hello", name)
}
