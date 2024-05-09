package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://hiteshchoudhary.com/"

func main() {
	fmt.Println("This is a web request.")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close() //This statement is very important responsibility

	data_bytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(data_bytes)
	fmt.Println(content)

}
