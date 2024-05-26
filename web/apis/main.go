package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Model schema for course
type Course struct {
	CourseId    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// dummy db
var courses []Course

// middleware or helper methods
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

//controllers

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>APIs in \"go\" go burrrrrrrr.</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses route")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//get id from the request
	params := mux.Vars(r)

	//search the course id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode(struct {
		Message  string `json:"message"`
		CourseID string `json:"course_id"`
	}{
		Message:  "Course not found",
		CourseID: params["id"],
	})
}

func main() {
	fmt.Println("This is a api.")
}
