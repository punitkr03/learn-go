package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name      string `json:"coursename"`
	Price     int    `json:"price"`
	Platform  string
	Passwrord string   `json:"-"`
	Tags      []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("JSON hell is in")
	// encodeJson()
	DecodeJson()
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

func DecodeJson() {
	sampleJsonData := []byte(`
	{
		"coursename": "AngularJS",
		"price": 483,
		"Platform": "YT",
		"tags": [
			"gulag",
			"dev",
			"js"
		]
	}`)

	var testCourse course

	isValid := json.Valid(sampleJsonData)

	if isValid {
		fmt.Println("JSON data here")
		json.Unmarshal(sampleJsonData, &testCourse)
		fmt.Printf("%#v\n", testCourse)
	} else {
		fmt.Println("Invalid json")
	}

	//some cases where we just want to add some data to key value pair

	var myData map[string]interface{}
	json.Unmarshal(sampleJsonData, &myData)
	fmt.Printf("%#v\n", myData)
}
