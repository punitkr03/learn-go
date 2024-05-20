package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name      string `json:"coursename"`
	Price     int    `json:"price"`
	Platform  string
	Passwrord string `json:"-"`
	Tags      []string `json:"tags,omitempty"`
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

	//package this data as json data

	finalJson, err := json.MarshalIndent(demoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalJson))

}
