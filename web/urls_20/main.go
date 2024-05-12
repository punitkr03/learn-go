package main

import (
	"fmt"
	"net/url"
	"strings"
)

const myurl = "https://localhost:3000/learn?coursename=react&payment_id=hello&payment_id=skfjh"

func main() {
	fmt.Println("Handling URLs in GoLang.")
	fmt.Println(myurl)

	//parsing
	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	fmt.Println(strings.Split(result.RawQuery, "&"))
}
