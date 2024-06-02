package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type course struct {
	ID       int    `json:"courseid"`
	Name     string `json:"coursename"`
	Price    int    `json:"price"`
	Platform string `json:"platform"`
}

// fake db
var sampleCourses = []course{
	{1, "ReactJS", 433, "YT"},
	{2, "MERN", 333, "YT"},
	{3, "AngularJS", 483, "YT"},
	{4, "Perl", 673, "YT"},
	{5, "Go", 543, "YT"},
	{6, "Python", 623, "YT"},
	{7, "Java", 723, "YT"},
	{8, "C++", 823, "YT"},
	{9, "Ruby", 923, "YT"},
	{10, "PHP", 823, "YT"},
}

func main() {
	fmt.Println("Hello web")
	greet()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/getonecourse/{courseid}", getOneCourse).Methods("GET")
	r.HandleFunc("/getallcourses", getAllCourses).Methods("GET")
	log.Fatal(http.ListenAndServe(":6969", r))
}

func greet() {
	fmt.Println("Hello from greet")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello from router </h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all the courses.")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sampleCourses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course route.")
	w.Header().Set("Content-Type", "application/json")

	//get the req params from the req
	params := mux.Vars(r)
	fmt.Println(params)

	//loop and find the matching id and find the response
	for _, course := range sampleCourses {
		courseid, err := strconv.Atoi(params["courseid"])
		if err != nil {
			// handle the error
			json.NewEncoder(w).Encode(struct {
				Message string `json:"message"`
			}{
				Message: "Invalid or no course id received.",
			})
			return
		}
		if course.ID == courseid {
			json.NewEncoder(w).Encode(course)
			return
		} else {
			json.NewEncoder(w).Encode(struct {
				Message string `json:"message"`
				ID      int    `json:"courseid"`
			}{
				Message: "No course found with the given id.",
				ID:      course.ID,
			})
		}
	}
}
