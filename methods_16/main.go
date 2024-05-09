package main

import "fmt"

func main() {
	fmt.Println("Methods in golang!!")
	//inheritance do not exist in golang

	user_1 := User{"Punit", "abc@abc.com", true, 20}
	fmt.Printf("User details are %+v\n", user_1)
	user_1.GetStatus()
	user_1.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active; ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@test.com"
	//This makes a copy and passes the value.
	fmt.Println("Email of this user is ", u.Email)
}
