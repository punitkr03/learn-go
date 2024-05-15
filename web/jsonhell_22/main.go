package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name      string
	Price     int
	Platform  string
	Passwrord string
	Tags      []string
}

func main() {
	fmt.Println("JSON hell is in")
	encodeJson()
}

func encodeJson() {
	demoCourses := []course{
		{"ReactJS", 433, "YT", "1234", []string{"web", "dev", "js"}},
		{"MERN", 333, "YT", "1234", []string{"web", "dev", "js"}},
		{"AngularJS", 483, "YT", "1234", []string{"gulag", "dev", "js"}},
		{"Perl", 673, "YT", "1234", nil},
	}

	//pakage this data as json data

	finalJson, _ := json.Marshal(demoCourses)

	fmt.Println(string(finalJson))

}
