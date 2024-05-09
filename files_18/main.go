package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files in go.")
	content := "This is a text to be sent in a file."

	file, err := os.Create("./testgofile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("Length of the file is :", length)
	defer file.Close()
	readFile("./testgofile.txt")
}

func readFile(filepath string) {
	data_bytes, err := os.ReadFile(filepath) //returns data in bytes

	if err != nil {
		panic(err)
	}

	fmt.Println("Text data inside the file is :\n", string(data_bytes))
}
