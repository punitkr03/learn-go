package main

import "fmt"

func main() {
	fmt.Println("Structs in golang!!")
	//inheritance do not exist in golang

	user_1 := User{"Punit", "abc@abc.com", true, 20}
	fmt.Printf("User details are %+v\n", user_1)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
